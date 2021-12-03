require 'benchmark'
require 'pry'

class Measure
  def initialize(name, func)
    @func = func
    @result = nil
    @name = name
  end

  def run
    time = Benchmark.measure do
      @result = @func.run
    end
    puts "#{@name} ANSWER: #{@result}, finished in #{time.real.round(2)}"
  end
end
