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
  - **go mod init <ชื่อ module> หรือ <ชื่อ project> หรือ แนะนำเป็นชื่อ root directory ของ project**
    - ex. `go mod init baby_go`

&nbsp;\
4. *clone repo มายัง path ที่สร้างด้วยคำสั่ง go mod init ... จะได้ไฟล์ __go.sum__*
  - **go get -u github.com/mephis5150/<repo\>**
    - ex. `go get -u github.com/mephis5150/baby_go`

&nbsp;\
5. *สามารถ import package และใช้งานได้แล้ว*
