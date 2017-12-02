let world = new World(document.body)

world.Run()

let socket = new WebSocket("ws://127.0.0.1:8080/socket")

socket.onmessage = function (evt) {
    let data = JSON.parse(evt.data)
    let type = data.type, message = data.message
    switch (type) {
        case "ACTOR_CREATE":
            world.CreateActor(message.id, message.pos, message.rot)
            break
        case "ACTOR_UPDATE":
            world.UpdateActor(message.id, message.pos, message.rot)
            break
        case "ACTOR_UPDATE_SHAPE_CIRCLE":
            world.UpdateActorShapeCircle(message.id, message.radius)
            break
        case "ACTOR_UPDATE_SHAPE_CONVEX":
            world.UpdateActorShapeConvex(message.id, message.points)
            break
        default:
            console.log(data)
            break
    }
}
