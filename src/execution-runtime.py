import sys
from lynx.common.scene import Scene
from lynx.common.object import Object
import json

input = input()
scene = Scene.deserialize(input)
print("{\"types\":[\"Floor\"],\"objects\":[\"{\\\"id\\\":172,\\\"position\\\":{\\\"x\\\":0,\\\"y\\\":0}}\"]}")
