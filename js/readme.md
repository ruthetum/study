# JavaScript

## Babel, Webpack, Polyfill
### Babel
- **JS 컴파일러**로 새로운 방식의 자바스크립트로 개발 후, 배포할 때에는 **예전 방식의 JS로 변환해서 배포**하기 위해 사용
- why?
  - 최신 버젼의 자바스크립트가 실행이 안되는 구버젼 웹브라우저를 대응하기 위해 
  - ES6 코드를 ES5 코드로 변환해주는 일에서 리액트의 JSX문법, 타입스크립트, 코드 압축, Proposal 까지 처리

### Webpack
- 모듈을 번들시켜주는 역할
- **빌드**를 통해 하나의 파일로 생성 
  - 빌드란 소스코드 파일을 실행가능한 소프트웨어 산출물로 만드는 과정으로 컴파일, 배포 등의 과정이 존재
- 바벨을 사용하려면 **Node.js**가 되어있어야 하고 터미널에서 웹팩 관련된 것들을 설치

### Polyfill
- 최신 **ECMAScript 환경** 생성
- 바벨은 ES6 => ES5로 변환할 수 있는 것들만 변환
  - ES6에서 비동기 처리를 위해 등장한 Promise와 같이 ES5에서 변환할 수 있는 대상이 없는 경우 에러가 발생한
  - 이럴 때 Polyfill을 이용해서 이슈를 해결할 수 있음
- example
    ```javascript
    const path = require('path'); // path 라이브러리를 가져온다
    const webpack = require('webpack');
    
    module.exports = {
        mode: 'development', // 개발용이 dev, 실서비스는 production으로 변경
        devtool: 'eval', // 빠르게
        resolve: { // entry app의 변환할 파일명
            extensions: ['.jsx', '.js'],
        },
        entry: { // 코드의 시작지점 (입력)
            app: './src/app.js' // 여러개면 배열로 넣으면 됨
        },
        module: { // 웹팩이 사용할 플러그인 지정
            rules: [{ // 여러개의 규칙들 (배열)
                test: /\.jsx?/, // 규칙 적용할 대상 확장자 (정규 표현식)
                                // jsx? => js, jsx
                exclude: /node_modules/, // 제외
                loader: 'babel-loader',
                options: {
                    presets: [ // plugin 설정들의 모음
                        ['@babel/preset-env', {
                            targets: { // 예전 브라우저 지원
                                browsers: ['> 1% in KR'], 
                            }, // 한국에서 1% 이상 점유율 가진 브라우저
                            debug: true, // 개발용
                        }], 
                        '@babel/preset-react'],
                    plugins: [],
                },
            }],
        },
        plugins: [
            new webpack.LoaderOptionsPlugin({debug:true}),
        ],
        output: { // 컴파일한 코드를 내놓을 위치 (출력)
            path: path.join(__dirname, 'dist'), // 파일위치할 디렉토리를 절대 경로로 지정
            filename: 'app.js' // 저장할 파일명 지정
        },
    }
    ```

#### ref. 
- https://yamoo9.gitbook.io/webpack/react/create-your-own-react-app/configure-polyfills
- https://iancoding.tistory.com/175
- http://yoonbumtae.com/?p=1140
---