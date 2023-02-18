import sys
import json

input = input()
print(f"Input: {input}", file=sys.stderr)
scene = json.loads(input)
print("{\"types\":[\"Floor\"],\"objects\":[\"{\\\"id\\\":172,\\\"position\\\":{\\\"x\\\":0,\\\"y\\\":0}}\"]}")
