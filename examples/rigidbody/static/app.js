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
        let camera = new THREE.PerspectiveCamera(75, window.innerWidth/window.innerHeight, 0.1, 1000)
        camera.position.z = 50

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

        actor.cube.position.x = position.x
        actor.cube.position.y = position.y
        actor.cube.position.z = position.z
    }

    has(Id) {
        return this.actors.has(Id)
    }
}

function CreateThreePosition(position) {
    return new THREE.Vector3(position.Y, position.Z, -position.X)
}

class Actor {
    constructor(id, position) {
        this.Id = id
        let geometry = new THREE.BoxGeometry( 1, 1, 1 )
        let material = new THREE.MeshBasicMaterial( { color: 0x00ff00 } )
        this.cube = new THREE.Mesh( geometry, material )

        this.cube.position.x = position.x
        this.cube.position.y = position.y
        this.cube.position.z = position.z
    }
}

let world = new World()

new Socket((name, json) => {
    if (world.has(json.Id)) {
        world.updatePosition(json.Id, CreateThreePosition(json.Position))
    } else {
        let actor = new Actor(json.Id, CreateThreePosition(json.Position))
        world.add(actor)
    }
})
