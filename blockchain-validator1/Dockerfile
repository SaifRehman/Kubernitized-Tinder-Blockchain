FROM node:carbon
WORKDIR /usr/src/app
COPY package*.json ./
COPY privkey0.json ./
COPY .env-node1 ./
RUN npm install
COPY . .
EXPOSE 30090 30091
CMD [ "node", "node1.js" ]
