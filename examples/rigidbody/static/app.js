class Socket {
    constructor(handle) {
        this.socket = new WebSocket("ws://" + window.location.host + "/socket")
        this.socket.onmessage = (event) => {
            const data = JSON.parse(event.data)

            handle(data.Name, JSON.parse(data.Json))
        }
    }
}

class World {
    constructor() {
        this.scene = new THREE.Scene()
        this.actors = new Map()
        let camera = new THREE.OrthographicCamera(window.innerWidth / -2, window.innerWidth / 2, window.innerHeight / 2, window.innerHeight / -2, -500, 1000 );

        let renderer = new THREE.WebGLRenderer()
        renderer.setSize(window.innerWidth, window.innerHeight)

        document.body.appendChild(renderer.domElement)

        let animate = () => {
            requestAnimationFrame(animate)

            renderer.render(this.scene, camera)
        }

        animate()
    }

    add(actor) {
        this.actors.set(actor.Id, actor)
        this.scene.add(actor.cube)
    }

    updatePosition(id, position) {
        let actor = this.actors.get(id)

        actor.cube.position.x = position.X
        actor.cube.position.y = position.Y
    }

    has(Id) {
        return this.actors.has(Id)
    }
}

class Actor {
    constructor(id, position) {
        this.Id = id
        let geometry = new THREE.BoxGeometry( 10, 10, 10 )
        let material = new THREE.MeshBasicMaterial( { color: 0x00ff00 } )
        this.cube = new THREE.Mesh( geometry, material )

        this.cube.position.x = position.X
        this.cube.position.y = position.Y
    }
}

let world = new World()

new Socket((name, json) => {
    if (world.has(json.Id)) {
        world.updatePosition(json.Id, json.Position)
    } else {
        let actor = new Actor(json.Id, json.Position)
        world.add(actor)
    }
})
