class World {
    constructor(body) {
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
        }

        requestAnimationFrame(update)
    }

    CreateActor(id, pos, rot) {
        let actor = new PIXI.Graphics()
        actor.x = pos.x
        actor.y = pos.y
        actor.rotation = rot * PIXI.DEG_TO_RAD
        actor.lineStyle(1, 0xFFFFFF, 1)

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

    DeleteActor(id) {
        let actor = this.actors[id]
        this.actors[id] = null
        this.stage.removeChild(actor)
    }
}
