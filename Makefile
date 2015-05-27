all: jose-validate JwsTest.class

jose-validate: jose-validate.go
	go build jose-validate.go

JwsTest.class: JwsTest.java
	javac -classpath jose4j-0.4.2.jar JwsTest.java

test: jose-validate JwsTest.class
	java -classpath jose4j-0.4.2.jar:commons-logging-1.2.jar:. JwsTest \
		2>/dev/null \
		| ./jose-validate

clean:
	if [ -e jose-validate ]; then rm jose-validate; fi
	if [ -e JwsTest.class ]; then rm JwsTest.class; fi
