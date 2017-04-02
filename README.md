# tracksy_backend

```
let ws = new WebSocket('ws://localhost:4001');

let message= {
  name: 'EARLY_ACCESS',
  data: {
    email: 'nathan1q@jsbin1.com'
  }
}

ws.onopen = () => {
  ws.send(JSON.stringify(message))
}

ws.onmessage = (e) => {
  console.log(JSON.parse(e.data));
}
```