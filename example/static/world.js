class World {
    constructor(body) {
        this.selectedActorId = null

        const WIDTH = window.innerWidth, HEIGHT = window.innerHeight
        this.renderer = PIXI.autoDetectRenderer(WIDTH, HEIGHT)
        body.appendChild(this.renderer.view)

        this.stage = new PIXI.Container()
        this.stage.position.set(WIDTH / 2, HEIGHT / 2)
        this.stage.scale.set(1, -1)
        this.actors = new Map()
    }

    Run() {
        let update = () => {
            requestAnimationFrame(update)

            this.renderer.render(this.stage)

            if ( this.selectedActorId ) {
                socket.send(JSON.stringify(new ActorCreate(
                    this.selectedActorId,
                    this.renderer.plugins.interaction.mouse.getLocalPosition(this.stage)
                    )))
            }
        }

        requestAnimationFrame(update)
    }

    CreateActor(id, pos, rot) {
        let actor = new PIXI.Graphics()
        actor.x = pos.x
        actor.y = pos.y
        actor.rotation = rot * PIXI.DEG_TO_RAD
        actor.lineStyle(1, 0xFFFFFF, 1)
        actor.interactive = true;
        actor.buttonMode = true;
        actor.hitArea = new PIXI.Circle(0, 0, 100);
        actor.on('pointerdown', () => {
            this.selectedActorId = id
        })
        actor.on('pointerup', () => {
            this.selectedActorId = null
        })
        actor.on('pointerout', () => {
            this.selectedActorId = null
        })

        this.actors.set(id, actor)
        this.stage.addChild(actor)

        return actor
    }

    UpdateActor(id, pos, rot) {
        let actor = this.actors.get(id)
        actor.x = pos.x
        actor.y = pos.y
        actor.rotation = rot * PIXI.DEG_TO_RAD
    }

    UpdateActorShapeCircle(id, radius) {
        let actor = this.actors.get(id)
        actor.drawCircle(0, 0, radius)
    }

    UpdateActorShapeConvex(id, points) {
        let actor = this.actors.get(id)

        let pixiPoints = []
        for (let p of points) {
            pixiPoints.push(new PIXI.Point(p.x, p.y))
        }
        let first = points[0]
        pixiPoints.push(new PIXI.Point(first.x, first.y))

        actor.drawPolygon(pixiPoints)
    }

    DebugLineCreate(start, end) {
        let line = new PIXI.Graphics()
        line.lineStyle(1, 0x00FF00, 1)

        line.moveTo(start.x, start.y)
        line.lineTo(end.x, end.y)

        this.stage.addChild(line)
        setTimeout(()=>{
            this.stage.removeChild(line)
        }, 1000)
    }

    DeleteActor(id) {
        let actor = this.actors[id]
        this.actors[id] = null
        this.stage.removeChild(actor)
    }
}
