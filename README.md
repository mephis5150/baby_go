### Dowload package แบบ manual (Old!)

<p>&nbsp;</p>
1. *ค้นหา go environment path*
```
go env GOPATH
```

<p>&nbsp;</p>
2. *เข้าไปยัง path และสร้าง directory __src__ ซึ่งจะใช้เก็บไฟล์ source code ของโปรเจ็ค*

<p>&nbsp;</p>
3. *สร้างตัวจัดการ package (คล้าย ๆ package.json ของ Node js)*
  - **go mod init <** ~~GOPATH/pkg/mod/~~ **DIR>/\<sub DIR>**
    - ex. `go mod init github.com/mephis5150`

<p>&nbsp;</p>
4. *clone repo มายัง path ที่สร้างด้วยคำสั่ง go mod init ...*
  - **go get -u github.com/mephis5150/\<repo>/\<package>**
    - ex. `go get -u github.com/mephis5150/baby_go/hello`

<p>&nbsp;</p>
5. *สามารถ import package และใช้งานได้แล้ว*
