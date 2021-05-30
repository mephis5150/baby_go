### Dowload package โดยใช้ go mod, go get

&nbsp;
1. *ค้นหา go environment path*
```
go env GOPATH
```

&nbsp;\
2. *เข้าไปยัง path และสร้าง directory __src__ ซึ่งจะใช้เก็บไฟล์ source code ของโปรเจ็ค*

&nbsp;\
3. *สร้างตัวจัดการ package (คล้าย ๆ package.json ของ Node js) จะได้ไฟล์ __go.mod__*
  - **go mod init <** ~~GOPATH/pkg/mod/~~ **DIR>/<sub DIR\>** หรือ **ชื่อ module**
    - ex. `go mod init github.com/mephis5150`
    - ex. `go mod init BabyGo`

&nbsp;\
4. *clone repo มายัง path ที่สร้างด้วยคำสั่ง go mod init ... จะได้ไฟล์ __go.sum__*
  - **go get -u github.com/mephis5150/<repo\>**
    - ex. `go get -u github.com/mephis5150/baby_go`

&nbsp;\
5. *สามารถ import package และใช้งานได้แล้ว*
