# Ryazan Tourist Route

___

# Overview

`back` backend service  
`front` frontend service

___
- I'm not that good in frontend, so idk how to pass envs from
docker-compose to vue app.  
- So, yeah, there is a docker-compose file, but base url for web app is hardcoded and https://rtr.hipahopa.ru/api.  
- To change url it's needed to rebuild docker image. You have to change line 9 in `front/src/App.vue` beforehand.