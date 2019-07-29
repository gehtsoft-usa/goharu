@echo off
echo to be done:
findstr /r /n "^\-" libharu.txt | find /c ":"
echo done:
findstr /r /n "^\+" libharu.txt | find /c ":"
echo question:
findstr /r /n "^\?" libharu.txt | find /c ":"