CodeRunner
==========

#### Server for running code. It is a micro system for supply result of running code.
#### This micro server can be use in entire system of developer community, education tool of programmingand so on.

#### 프로그램 실행 서버입니다. HTTP통신을 통하여 코드를 입력으로 받아, 해당 코드를 실행하고 결과값을 돌려주는 Task를 담당합니다. 고객의 코드의 실행이 필요한 교육 도구, 개발자 커뮤니티 등 웹 기반의 프로젝트의 서버와 HTTP 통신을 통하여 코드를 실행하고 결과값을 줄 것입니다.

#### 고루틴을 통하여 병렬 수행을 하고, 컨테이너로써 배포하여 확장성을 지닙니다. 다수의 컨테이너로 배포할 경우 컨테이너 오케스트레이션 도구를 사용하시는 것이 좋습니다.






1. Installation
----------------

* For easy installation, please use docker
* If you do not want to install go, just execute main by "./main"






2. Build
--------


* By using docker
```
make docker V={put tag}
docker run -i -p 3001:3001 runserver:{tag}
```



* By using go build

```
make build
```






3. Usage
--------

Send post request to under point
```
http://{hostname}:3001/code
or
https://{hostname}:3001/code
```


```
POST json {
	Text string,
	Filename string
	Language string
}
```


Now only support java, but will support cplusplus, python, javascript.
and will be support Language Version





### Supported Language
```
java
c++
python
```



made by suhwanggyu
