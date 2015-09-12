# BaDNS

Mess up with DNS responses.

```
$ badns -p 5555 --host-prefix www --replace-type A:MX --delay 5 &


$ dig -p 5555 @localhost www.google.com

; <<>> DiG 9.9.5-3ubuntu0.5-Ubuntu <<>> -p 5555 @localhost www.google.com
; (1 server found)
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 45059
;; flags: qr rd ra; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 1

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 512
;; QUESTION SECTION:
;www.google.com.                        IN      A

;; ANSWER SECTION:
www.google.com.         256     IN      MX      55354 google.com.

;; Query time: 3027 msec
;; SERVER: 127.0.0.1#5555(127.0.0.1)
;; WHEN: Sat Sep 12 00:34:11 PDT 2015
;; MSG SIZE  rcvd: 73


$ dig -p 5555 @localhost www.metin.org
;; Got bad packet: bad compression pointer
103 bytes
dc 27 81 80 00 01 00 02 00 00 00 01 03 77 77 77          .'...........www
05 6d 65 74 69 6e 03 6f 72 67 00 00 01 00 01 03          .metin.org......
77 77 77 05 6d 65 74 69 6e 03 6f 72 67 00 00 05          www.metin.org...
00 01 00 00 0d f3 00 0b 05 6d 65 74 69 6e 03 6f          .........metin.o
72 67 00 05 6d 65 74 69 6e 03 6f 72 67 00 00 0f          rg..metin.org...
00 01 00 00 0d f3 00 04 ad ff ea 94 00 00 29 02          ..............).
00 00 00 00 00 00 00

```
