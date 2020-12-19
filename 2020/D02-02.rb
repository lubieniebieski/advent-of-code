require_relative 'lib/utils.rb'

class Calculation
  def initialize
    @passwords = IO.readlines('./D02-input')
  end

  def run
    @passwords.find_all do |p|
      capture = p.match(/(?<yes_index>\d+)-(?<no_index>\d+)\s(?<letter>\D):\s(?<string>\S+)/)

      (capture['string'].chars[capture['yes_index'].to_i - 1] == capture['letter']) ^
        (capture['string'].chars[capture['no_index'].to_i - 1] == capture['letter'])
    end.size
  end
end

Measure.new('D02-02', Calculation.new).run
