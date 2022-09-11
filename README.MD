# justlog [![Build Status](https://github.com/gempir/justlog/workflows/ci/badge.svg)](https://github.com/gempir/justlog/actions?query=workflow%3Aci)

### ¿Qué es esto?

Justlog es un bot IRC de Twitch. Se centra en el registro de mensajes y proporcionar una API para los registros.

### API

La documentación de la API se puede ver a través del frontend de justlog haciendo clic en el símbolo "docs" aquí:  
![image](https://user-images.githubusercontent.com/1629196/159481078-0de98f01-2816-49bd-8e17-ba7cf66cb064.png)

### Docker

```
mkdir logs
docker run -p 8025:8025 --restart=unless-stopped -v $PWD/config.json:/etc/justlog.json -v $PWD/logs:/logs ghcr.io/gempir/justlog
```

### Comandos

Sólo los administradores pueden utilizar estos comandos

- `!justlog status` Responderá con el tiempo de actividad
- `!justlog join gempir pajlada` Se unirá a los canales y los añadirá a la configuración
- `!justlog part gempir pajlada` Se retirará de los canales y los eliminará de la configuración
- `!justlog optout gempir gempbot` Excluirá a los usuarios del registro de mensajes o de la consulta de los registros anteriores de ese usuario, lo mismo se aplica al propio canal de los usuarios
- `!justlog optin gempir gempbot` Revertirá la exclusión de los registros

### Configuración

```
{
    "admins": ["gempir"], // Sólo responderá a los comandos ejecutados por estos usuarios
    "logsDirectory": "./logs", // El directorio en el que estarán los registros
    "adminAPIKey": "noshot", // Su API Key secreta para acceder a la API de admin, puede ser cualquier cadena de texto, es utilizada en la solicitud API a los endpoints de admin
    "username": "gempbot", // Usuario del bot (puede ser justinfan123123 si no quieres usar una cuenta)
    "oauth": "oauthtokenforchat", // Token del bot, puede ser cualquier cosa si el usuario es justinfan123123
    "botVerified": true, // Aumenta los ratelimits si tienes un bot verificado, el bot podrá unirse más rápido, falso por defecto
    "clientID": "mytwitchclientid", // Su Client ID, necesario para obtener UserIDs o nombres de usuario, etc
    "clientSecret": "mysecret", // Su Client Secret de Twitch
    "logLevel": "info", // El nivel de registro, deberías mantenerlo en info, las demás opciones son: trace, debug, info, warn, error, fatal, y panic, logs output es stdout
    "channels": ["77829817", "11148817"], // Los canales (UserIDs) que desea registrar
    "archive": true // Probablemente se mantenga en true, deshabilitará el gzipping de los registros antiguos si es false, útil si configura la compresión por su cuenta
}
```

### Desarrollo

El desarrollo requiere [yarn](https://classic.yarnpkg.com/) y [go-swagger](https://goswagger.io/)

Ejecuta `go build && ./justlog` y `yarn start` en la carpeta web. 

O ejecuta `make container` y `make run_container`
