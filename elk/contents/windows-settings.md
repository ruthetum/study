# 윈도우 실습 환경 구성

## 엘라스틱서치 설치
- 엘라스틱 가이드(https://www.elastic.co/guide/index.html) 접속

- 버전 확인 후 원하는 버전 클릭
    - LTS : https://www.elastic.co/guide/en/elastic-stack/current/index.html

- 설치 페이지 이동 후 Windows 버전 zip 파일 다운로드
    - LTS : https://www.elastic.co/guide/en/elasticsearch/reference/8.1/install-elasticsearch.html#install-elasticsearch

- 압축 해제 후 `/bin` 디렉토리 내에 `elasticsearch.bat` 실행

- 명령 프롬프트에서 `curl -X GET localhost:9200/?pretty` 명령어로 정상 작동 확인
    - 기본 설정 포트 : 9200

    - `?pretty` : JSON 결과를 가독성 좋게 보겠다

    - 만약 `curl: (52) Empty reply from server` 오류가 발생한다면 `/config/elasticsearch.yml`에 `xpack.security.enabled` 옵션을 `false`로 설정

        - 참고 : https://stackoverflow.com/questions/35921195/curl-52-empty-reply-from-server-timeout-when-querying-elastiscsearch

    <br>

    ![image](https://user-images.githubusercontent.com/59307414/164432474-3e0bd6d1-7e61-4de7-b0ad-f86c12696d53.png)


## 키바나 설치
- 설치 페이지 이동 후 Windows 버전 zip 파일 다운로드
    - LTS : https://www.elastic.co/guide/en/kibana/8.1/install.html#_install_kibana_yourself

- 압축 해제 후 `/bin` 디렉토리 내에 `kibana.bat` 실행

- localhost:5601 확인
    - 기본 설정 포트 : 5601
    
    - 엘라스틱서치가 먼저 실행되지 않으면 에러 발생

    - 만약 실행되지 않고 명령 프롬프트 창이 알아서 닫힌다면 `/config/kibana.yml`에서 `server.port`, `server.host`, `elasticsearch.hosts` 주석 해제

    


### 참고
- 엘라스틱서치 설치 : https://www.elastic.co/guide/en/elasticsearch/reference/current/install-elasticsearch.html