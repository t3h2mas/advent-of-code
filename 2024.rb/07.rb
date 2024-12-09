input = <<~TEXT
190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
TEXT

equations = input.split("\n").map do |line| 
  target, test_inputs = line.split(': ')
  { target: target.to_i, test_inputs: test_inputs.split(' ').map(&:to_i) }
end

# just build them all
# input 1,2,3
# output: [[1 + 2 + 3], [1 * 2 + 3], [1 + 2 * 3], [1 * 2 * 3]]
def build_evaluations(xs)
  root = Value.new(value: xs.first, edges: [])

  build_evaluation_steps(root, xs[1..])

  root
end

def build_evaluation_steps(parent, xs)
  return if xs.empty?

  value = xs.first

  new_sum = Value.new(value:, edges: [])
  new_mul = Value.new(value:, edges: [])

  parent.edges << Edge.new(operator: :+, operand: new_sum)
  parent.edges << Edge.new(operator: :*, operand: new_mul)

  build_evaluation_steps(new_sum, xs[1..])
  build_evaluation_steps(new_mul, xs[1..])
end

Value = Data.define(:value, :edges)

Edge = Data.define(:operator, :operand)

Value.new(value: 1, edges: [
  Edge.new(operator: :+, operand: Value.new(value: 2, edges: [])),
  Edge.new(operator: :*, operand: Value.new(value: 2, edges: []))
])

build_evaluations([1,2,3])

