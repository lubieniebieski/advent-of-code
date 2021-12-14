require 'pry'

class D14
  def initialize(data)
    data = data.map(&:strip)
    @code = {}
    @template = data.shift
    data.each do |line|
      if (match = line.match(/(?<pair>\S{2}) -> (?<element>\S)/))
        @code[match[:pair]] = match[:element]
      end
    end
  end

  def self.run(data)
    [new(data).part1, new(data).part2]
  end

  def part1
    step(10)

    occurences = calculate_occurences(@template)
    occurences.values.max - occurences.values.min
  end

  def part2
    # step(40)

    # occurences = calculate_occurences(@template)
    # occurences.values.max - occurences.values.min
  end

  def step(number_of_steps)
    number_of_steps.times do |_time|
      @template =
        @template.chars.each_cons(2).map do |pair|
          [pair.first, decode_pattern(pair.join) || pair.last]
        end.append(@template[-1]).join
    end
    @template
  end

  def decode_pattern(pattern)
    @code[pattern] ||= pattern.chars.each_cons(2).map do |pair|
      [pair.first, decode_pattern(pair.join) || pair.last]
    end.append(pattern[-1]).join
  end

  def calculate_occurences(template)
    occurences = {}
    template.chars.uniq.each do |char|
      occurences[char] = template.count(char)
    end
    occurences
  end
end
