from lynx.common.scene import Scene
from lynx.common.actions.move import Move
from lynx.common.vector import Vector

input = input()
scene = Scene.deserialize(input)
scene.add_entity(Move(object_id=123, vector=Vector(1,1)))
print(scene.serialize())
