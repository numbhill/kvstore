# Key-Value Store Go

Ky projekt është një shembull i një sistemi të shpërndarë klient-server në gjuhën Go, që implementon një shërbim storage në formën e një key-value store (çelës-vlerë) duke përdorur Remote Procedure Call (RPC). Ky README do të ju udhëzojë në krijimin dhe ekzekutimin e projektin.

## Server

Kodi në server.go është pjesa e serverit. Ky kod krijon një server TCP në portin 1234 dhe ofron dy operacione për manipulimin e të dhënave në storage: Put dhe Get.

## Client

Kodi në client.go është pjesa e klientit. Ky kod është përgjegjës për dërgimin e kërkesave RPC për të manipuluar dhe marrë të dhëna nga serveri.

### Hapat për të ekzekutuar projektin

1. **Krijoni strukturën e skedarëve siç është treguar më lartë.**

2. **Vendosni kodin e dhënë në skedarët përkatës (server.go dhe client.go).**

3. **Në terminal, navigoni te direktoria server dhe ekzekutoni komandën:**

    ```bash
    go run server.go
    ```

4. **Në një terminal tjetër, navigoni te direktoria client dhe ekzekutoni komandat për të manipuluar të dhënat:**

    ```bash
    go run client.go put "1" "test"
    go run client.go get "1"
    ```

## Struktura e projektit

```plaintext
kvstore/
├── server/
│   │   └── server.go
├── client/
│   │   └── client.go
├── storage.json
```
