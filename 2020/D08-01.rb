require_relative 'lib/utils.rb'

class Calculation
  def initialize
    @instructions = IO.read('./D08-input').split("\n").map { |i| Instruction.new(i) }
    @index = 0
    @accumulator = 0
  end

  def run
    loop do
      instruction = @instructions[@index]
      return @accumulator if instruction.visited
      process_instruction(instruction)
    end
  end

  private

  def process_instruction(instruction)
    instruction.visit!
    if instruction.jmp?
      increase_index!(instruction.value)
    else
      increase_index!(1)
      increase_accumulator!(instruction.value) if instruction.acc?
    end
    instruction.visit!
  end

  def increase_index!(value = 1)
    @index += value
  end

  def increase_accumulator!(value)
    @accumulator += value
  end
end

class Instruction
  attr_reader :instruction_type, :value, :visited

  def initialize(data)
    @instruction_type, value = data.split(' ')
    @value = value.to_i
    @visited = false
  end

  def visit!
    @visited = true
  end

  def jmp?
    instruction_type == 'jmp'
  end

  def nop?
    instruction_type == 'nop'
  end

  def acc?
    instruction_type == 'acc'
  end
end

Measure.new('D08-01', Calculation.new).run
