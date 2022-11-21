### aas_et generator

##### Examples

```
go install
go build aas_et.go

./aas_et your-username your-password [output-file]
```

#### Example response output file

```json
{
  "SID": "BAD_COOKIE",
  "LSID": "BAD_COOKIE",
  "Token": "aas_et/YJSNPI_1145141919810",
  "Services": "hist,mail,lso,talk,chromiumsync,android,cl,sj,lh2,cloudprint,ah,ahadmin,youtube,friendview,multilogin,chromeoslogin,writely,omaha,androidconsole,mymaps,googleplay,sitemaps,gtrans,local,analytics,grandcentral,sierra,memento,groups2,nova,uif,domains,chromewebstore,billing,googleone,wise,g1phonebackup"
}
```

#### How to use in 2FA mode

You can generate a APP password in Google account settings.

https://myaccount.google.com/apppasswords