### First Personal Project

Starting code of my first personal project for Boot.dev (well, *maybe*, *maybe not*) ...at the moment I plan for it to be a CLI tool for analyzing Wireshark CSV exports...

Example console output after running...

Read 33993 bytes from file
```bash
[No. Time Source Destination Protocol Length Info]
[1 0.000000 192.168.0.173 192.168.0.205 SSH 138 Client: Encrypted packet (len=84)]
[2 0.013521 192.168.0.205 192.168.0.173 TCP 60 22  >  20018 [ACK] Seq=1 Ack=85 Win=13007 Len=0]
[3 0.015089 192.168.0.173 192.168.0.205 SSH 138 Client: Encrypted packet (len=84)]
[4 0.026408 192.168.0.205 192.168.0.173 TCP 60 22  >  20018 [ACK] Seq=1 Ack=169 Win=13007 Len=0]
[5 1.234611 192.168.0.173 192.168.0.255 UDP 70 62976  >  22222 Len=28]
[6 1.235236 192.168.0.173 192.168.0.255 UDP 70 62983  >  22222 Len=28]
[7 1.235240 192.168.0.173 192.168.0.255 UDP 56 62990  >  3289 Len=14]
[8 1.757896 192.168.0.205 192.168.0.173 SSH 138 Server: Encrypted packet (len=84)]
[9 1.812759 192.168.0.173 192.168.0.205 TCP 54 20018  >  22 [ACK] Seq=169 Ack=85 Win=2045 Len=0]
[10 1.835559 192.168.0.173 149.102.224.175 WireGuard 138 Transport Data, receiver=0x2D55921C, counter=2, datalen=64]
[14 1.907417 149.102.224.175 192.168.0.173 WireGuard 122 Transport Data, receiver=0xCDCD3305, counter=2, datalen=48]
[15 2.018173 149.102.224.175 192.168.0.173 WireGuard 1482 Transport Data, receiver=0xCDCD3305, counter=3, datalen=1408]
[16 2.018173 149.102.224.175 192.168.0.173 WireGuard 202 Transport Data, receiver=0xCDCD3305, counter=4, datalen=128]
[17 2.018424 192.168.0.173 149.102.224.175 WireGuard 122 Transport Data, receiver=0x2D55921C, counter=5, datalen=48]
[18 2.019013 149.102.224.175 192.168.0.173 WireGuard 1482 Transport Data, receiver=0xCDCD3305, counter=5, datalen=1408]
[19 2.019013 149.102.224.175 192.168.0.173 WireGuard 1482 Transport Data, receiver=0xCDCD3305, counter=6, datalen=1408]
[20 2.019237 192.168.0.173 149.102.224.175 WireGuard 122 Transport Data, receiver=0x2D55921C, counter=6, datalen=48]
[21 2.052573 149.102.224.175 192.168.0.173 WireGuard 186 Transport Data, receiver=0xCDCD3305, counter=7, datalen=112]
# etc..........
[298 39.033723 192.168.0.173 149.102.224.175 WireGuard 122 Transport Data, receiver=0x2D55921C, counter=73, datalen=48]
[299 39.033895 192.168.0.173 149.102.224.175 WireGuard 122 Transport Data, receiver=0x2D55921C, counter=74, datalen=48]
[300 39.067192 149.102.224.175 192.168.0.173 WireGuard 122 Transport Data, receiver=0xCDCD3305, counter=78, datalen=48]
```
