# go-docker

Just to play around with docker

The goal of this project is to create a web server that updates itselves whenever a file has been changed.

The idea is to eliminate the manual steps required in the traditional workflow:

- Change code
- Stop running server
- Compile
- Test components
- Start server

The solution is to detect whenever code has been changed and automate the remaining tasks.

Note: the project depends on Gorilla, but I'll probably remove this dependency by my own functions a.s.a.p.
