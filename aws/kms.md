# AWS KMS (Key Management Service)

# AWS 암호화 서비스

<img width="645" alt="스크린샷 2023-04-18 오후 9 25 37" src="https://user-images.githubusercontent.com/59307414/232818231-6d276210-8cac-41ba-8994-f6c56d8572b4.png">

## ACM (AWS Cerificate Manager)

[https://aws.amazon.com/ko/certificate-manager/](https://aws.amazon.com/ko/certificate-manager/)

- 인증서를 발급하기 위해 사용하는 서비스
    - 인증서는 public한 공인 인증서 또는 private한 사설 인증서로 구분
- public 인증서는 AWS 서비스와 integration이 되어 있어 cloudfront, alb/nlb와 결합하여 사용 가능
- private 인증서는 aws native service로 활용하고 싶은 경우 활용


## KMS (Key  Management Service)

[https://aws.amazon.com/ko/kms/](https://aws.amazon.com/ko/kms/)

<img width="643" alt="스크린샷 2023-04-18 오후 9 34 25" src="https://user-images.githubusercontent.com/59307414/232818341-e4381577-9e6a-4fcb-a8a1-6c234e3ea5d3.png">

### 개요

- AWS 서비스 전체에 대해 암호화 키를 생성/관리/제어
    - 중앙집중형
- 사용 권한 부여 시 policy와 grant를 통해 권한 부여
    - 이 부분이 AWS IAM과 연동
- AWS의 API 기반으로 동작하기 때문에 관련 기록은 CloudTrail에 적재
- KMS의 이용 방식은 server side 암호화, client side 암호화 두 가지 방식으로 나눌 수 있음
    - AWS의 서비스에서 KMS를 사용하는 이유는 server side 암호화 목적
    - 사용자 또는 애플리키에션에서 사용하는 것은 client side 암호화를 이용
        - KMS의 키가 애플리케이션으로 전달되고, 클라이언트 측에서 해당 키를 통해 데이터를 암호화
- 목적에서 따라서 암호화 방식은 하나만 사용할 수도 있고, 둘 다 사용할 수도 있음

### 주요 개념

<img width="631" alt="스크린샷 2023-04-18 오후 9 44 25" src="https://user-images.githubusercontent.com/59307414/232818450-dd8c47a4-a6e5-4820-8ab0-f87876c7429b.png">

**CMK (**Customer Managed Key**)**

- 고객이 만드는 KMS Key
- KMS 내부에서 CMK를 이용하여 암호화/복호화 활용

**봉투암호화**

- 클라이언트 사이드에서 데이터 암호화 시에는 데이터 키를 사용하여 평문을 암호화 (이후 데이터키 삭제)
- 편지에 내용물을 담고, 주소/목적지를 적는 것처럼 데이터키로 암호화한 암호문과 함께 암호화된 데이터키를 함께 보관
- 이런 컨셉을 AWS에서는 봉투암호화로 표현

**Encryption Context**

- AWS KMS로 보호된 데이터와 연관시키고자 하는 부가정보(Key-value pair)
- 암호문과 별개로 암호문이 어떠한 유형의 문서인지, 어떤 내용을 포함하고 있는 메타데이터 (Condition 필드에 추가)
- 평문으로 저장됨

**Policy와 Grant**

- policy: CMK에 대한 정책
    - iam에서는 identity에 정책(id based policy)을 부여할 수도 있고, resource에 정책(resource based policy)을 부여할 수도 있음
    - KMS에서도 동일한 컨셉 적용, identity 기반으로 정책을 부여할 수도 있고, CMK에 대해서 정책을 부여할 수 있음
- grant: 일종의 토큰 개념
    - CMK에 대한 사용 권한을 identity 또는 CMK로 얻는 것이 아니라 grant(token)을 통해 부여받을 수 있음

### KMS 키의 계층 구조

<img width="602" alt="스크린샷 2023-04-18 오후 10 02 02" src="https://user-images.githubusercontent.com/59307414/232818513-32308b04-fbec-4da6-80b8-3b39ab12c727.png">

KMS 인프라는 내부적으로 **HSM**과 **고가용성이 보장된 스토리지 영역**이 존재

KMS를 사용하기 위해서는 CMK가 필요

- HSM을 통해 CMK를 생성 (= HSM이 KMS를 생성)
- CMK는 요청 시 HSM에 생성되고, **평문의 CMK**가 **HSM 영역의 메모리** 상에 적재

HSM 내부적으로는 **평문의 CMK**를 암호화하기 위한 별도의 **도메인 키** 존재

- 암호화 목적은 메모리에 있는 **평문의 CMK**를 메모리 상에서 암호화하기 위함이 아닌 CMK를 사용하지 않는 환경에서는 메모리에 CMK가 내려와야 되는데 저장소에 저장해야 함, 이 때 암호화하여 저장하기 위해 **도메인 키**를 사용 (암호화된 CMK를 스토리지에 저장)
- 이러한 원리로 인해 KMS의 CMK는 외부로 나갈 때 평문 형태로 유출되지 않음

그렇다면 외부의 클라이언트는 KMS 서비스를 어떻게 사용하는가?

- 클라이언트는 API 요청(GenerateDataKey)을 하면 KMS에서 **데이터 키**를 전달받게 됨
- KMS 내부적으로 데이터 키를 생성하고, **평문의 데이터 키**와 **CMK로 암호화한 데이터 키**를 클라이언트에게 쌍으로 전달

**계층적인 구조**

- KMS에서 발급해서 최종적으로 데이터를 암호화하는 키는 **데이터 키**
- 이 **데이터 키**의 보안을 위해 **CMK**가 사용
    - **데이터 키**는 **CMK**를 통해 암호화
- 이 **CMK**의 보호를 위해 **도메인 키** 사용
    - **CMK**는 **도메인 키**를 통해 암호화

### 데이터 키 생성 과정(대칭키)

<img width="660" alt="스크린샷 2023-04-18 오후 10 06 04" src="https://user-images.githubusercontent.com/59307414/232818695-66048715-e7b5-4ff6-9d94-92bef71636c1.png">

데이터 키를 요청하는 GenerateDataKey는 대칭키를 요청하는 API

데이터 암호화 시에는 데이터 키를 사용하여 암호화

- 더 이상 암호화할 데이터가 없거나 필요 없는 경우 데이터 키 삭제

이후 클라이언트는 데이터를 사용할 때 암호화된 데이터 키를 통해 데이터 복호화

이 경우에 데이터의 암호화는 **클라이언트 사이드에서 진행**

### KMS 에서의 암호화

<img width="660" alt="스크린샷 2023-04-18 오후 10 10 55" src="https://user-images.githubusercontent.com/59307414/232818746-18434160-ba81-40fa-b43f-3cf891b6728d.png">

데이터 키를 발급받아서 클라이언트 사이드에서 암호화를 제공하기 하지만 KMS를 통해서 **서버 사이드 암호화**를 할 수 있음 (Encrypt API)

### 비대칭키 사용

<img width="619" alt="스크린샷 2023-04-18 오후 10 38 19" src="https://user-images.githubusercontent.com/59307414/232818865-c710f5fa-6742-41d3-9dc5-fb488e072a53.png">

KMS를 통해 비대칭키도 발급받을 수 있음 (GenerateDataKeyPair)

- 전자서명 또는 암호화 활용 가능

<img width="666" alt="스크린샷 2023-04-18 오후 10 41 22" src="https://user-images.githubusercontent.com/59307414/232818944-c2453432-f1a8-4aa0-90e2-556aeda5e989.png">

대칭키의 경우 두 개의 키(평문의 데이터 키, 암호화된 데이터 키)가 반환됐던 반면, 비대칭키의 경우 세 개의 키(공개키, 비밀키, 암호화된 비밀키 by CMK)발급

**요약**

<img width="626" alt="스크린샷 2023-04-18 오후 11 20 09" src="https://user-images.githubusercontent.com/59307414/232819361-21345320-d598-422d-9022-05f970b14ef6.png">

### 봉투암호화 동작 원리

<img width="617" alt="스크린샷 2023-04-18 오후 11 04 25" src="https://user-images.githubusercontent.com/59307414/232819076-7ee243f1-ce9d-409b-b9a6-fcfb105779ae.png">

클라이언트 사이드에서 데이터 암호화 시에는 데이터 키를 사용하여 평문을 암호화

- 더 이상 암호화할 데이터가 없거나 필요 없는 경우 데이터 키 삭제
- 암호문으로 보관

이렇게 암호문으로 보관할 때 복호화 시에 사용할 암호화된 데이터키를 메타데이터로 함께 보관

- 편지에 내용물을 담고, 주소/목적지를 적는 것처럼 데이터키로 암호화한 암호문과 함께 암호화된 데이터키를 함께 보관
- 암호화에 사용한 데이터키는 삭제하기 때문

### **Encryption Context**

<img width="609" alt="스크린샷 2023-04-18 오후 11 11 39" src="https://user-images.githubusercontent.com/59307414/232819584-8e7ea7cc-bea3-4916-88c5-23b876f41225.png">

AWS KMS로 보호된 데이터와 연관시키고자 하는 부가정보(Key-value pair)

- 암호문과 별개로 암호문이 어떠한 유형의 문서인지, 어떤 내용을 포함하고 있는 메타데이터 (Condition 필드에 추가)
- 평문으로 저장

### Policy & Grants

<img width="646" alt="스크린샷 2023-04-18 오후 11 18 10" src="https://user-images.githubusercontent.com/59307414/232819963-236bd4d6-ec91-4a4d-9510-9e919da54399.png">

## CloudHSM (Hardware Security Module)

[https://aws.amazon.com/ko/cloudhsm/](https://aws.amazon.com/ko/cloudhsm/)

<img width="679" alt="스크린샷 2023-04-18 오후 11 35 58" src="https://user-images.githubusercontent.com/59307414/232820138-f4d61da8-b3a5-421f-8bba-6f8bd77bf947.png">

<img width="613" alt="스크린샷 2023-04-18 오후 11 30 54" src="https://user-images.githubusercontent.com/59307414/232820087-8af3e4a5-5e54-412a-9308-008457dfa1c2.png">

대체로 KMS와 기능은 유사하지만 CloudHSM이 필요한 경우가 존재

- 높은 수준의 인증이 요구되는 경우
    - KMS는 FIPS 140-2 Level 2 인증
- 단일 테넌트(single tenant) 기반
    - KMS는 멀티 테넌트(multi tenant) 기반
- 표준 기술 적용이 필요한 경우
    - KMS는 AWS API 기반으로 동작하기 때문에 표준 기술을 지원하지 않는 경우 있음
- 그 외 KMS에서 지원하지 않는 기능이 필요한 경우

<img width="662" alt="스크린샷 2023-04-18 오후 11 36 18" src="https://user-images.githubusercontent.com/59307414/232820277-2871c9db-99f8-4639-93e9-0a198e492b85.png">

CloudHSM은 IaaS 같은 서비스로 생각

- AWS가 관리하는 영역보다 사용자가 관리하는 영역의 범위가 큼

<img width="665" alt="스크린샷 2023-04-18 오후 11 37 25" src="https://user-images.githubusercontent.com/59307414/232820291-317cc851-6d5a-41da-b27d-f3d8f1faa396.png">

백업의 경우 사용자가 직접 관리하지 않아도 자동 백업 지원

- 이를 통해 다른 리전에서 복원도 가능
- 다만 교차 리전간 동기화는 지원하지 않음 (리전 별로 독립적으로 운영)

<img width="628" alt="스크린샷 2023-04-18 오후 11 38 43" src="https://user-images.githubusercontent.com/59307414/232820299-3fb2f79e-5732-4231-a3f1-6c0d53fd84c9.png">

KMS와 연계해서 사용도 가능

- KMS에서 발급한 CMK를 CloudHSM에 저장

# AWS 암호화 서비스 요약

<img width="693" alt="스크린샷 2023-04-18 오후 11 40 19" src="https://user-images.githubusercontent.com/59307414/232820306-b92531f8-6e6e-4d63-9808-553e8e9efde8.png">

# Reference

[https://aws.amazon.com/ko/kms/](https://aws.amazon.com/ko/kms/)

[https://docs.aws.amazon.com/ko_kr/kms/latest/developerguide/concepts.html](https://docs.aws.amazon.com/ko_kr/kms/latest/developerguide/concepts.html)

[https://docs.aws.amazon.com/ko_kr/kms/latest/developerguide/iam-policies.html](https://docs.aws.amazon.com/ko_kr/kms/latest/developerguide/iam-policies.html)