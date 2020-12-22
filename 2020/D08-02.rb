require_relative 'lib/utils.rb'

class Calculation
  def initialize
    @instructions = IO.read('./D08-input').split("\n").map { |i| Instruction.new(i) }
    @index = 0
    @accumulator = 0
  end

  def run
    @original_instructions = @instructions.clone.map(&:clone)
    @changing_index = 0
    loop do
      instruction = @instructions[@index]
      return @accumulator if instruction == @instructions.last
      process_visited && next if instruction.visited
      process_instruction(instruction)
    end
  end

  private

  def process_visited
    @index = 0
    @accumulator = 0
    @instructions, @changing_index =
      modify_instructions(@original_instructions.clone.map(&:clone), @changing_index)
  end

  def modify_instructions(instructions, index)
    instructions[index..-1].each_with_index do |i, ind|
      next if i.acc?
      i.switch_op!
      index = index + ind + 1
      break
    end

    [instructions, index]
  end

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

  def reset!
    @visited = false
  end

  def switch_op!
    @instruction_type = if instruction_type == 'jmp'
                          'nop'
                        else
                          'jmp'
                        end
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

Measure.new('D08-02', Calculation.new).run
