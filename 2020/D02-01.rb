require_relative 'lib/utils.rb'

class Calculation
  def initialize
    @passwords = IO.readlines('./D02-input')
  end

  def run
    @passwords.find_all do |p|
      capture = p.match(/(?<min>\d+)-(?<max>\d+)\s(?<letter>\D):\s(?<string>\S+)/)
      count = capture['string'].chars.count { |c| c == capture['letter'] }
      (capture['min'].to_i..capture['max'].to_i).cover?(count)
    end.size
  end
end

Measure.new('D02-01', Calculation.new).run
