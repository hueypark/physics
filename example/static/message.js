class ActorCreate {
    constructor(id, pos, rot) {
        this.type = "ACTOR_CREATE"
        this.id = id
        this.pos = pos
        this.rot = rot
    }
}
