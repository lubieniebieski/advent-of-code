require_relative '../solutions/D05'

describe D05 do
  context 'Real data' do
    let(:input_path) { File.join(File.dirname(__FILE__), '../inputs/D05.txt') }
    let(:input) { File.read(input_path).split("\n") }
    let(:subject) { described_class.new(input) }
    let(:part1) { subject.part1 }
    let(:part2) { subject.part2 }

    it 'Solves Part 1' do
      expect { part1 }.not_to raise_exception
      expect(part1).to be > 4636
    end

    it 'Solves Part 2' do
      expect { part2 }.not_to raise_exception
    end

    it 'Works with real data' do
      results = described_class.run(input)
      warn "\t#{results}"
      expect(results).not_to eq nil
    end
  end

  context 'Provided examples' do
    let(:test_data) do
      '0,9 -> 5,9
      8,0 -> 0,8
      9,4 -> 3,4
      2,2 -> 2,1
      7,0 -> 7,4
      6,4 -> 2,0
      0,9 -> 2,9
      3,4 -> 1,4
      0,0 -> 8,8
      5,5 -> 8,2'
    end
    let(:test_input) { test_data.split("\n").map(&:strip) }
    let(:subject) { described_class.new(test_input) }
    describe '.part1' do
      it 'Passes the example' do
        expect(subject.part1).to eq 5
      end
    end

    context 'Part 2' do
      it 'Passes the example' do
        expect(subject.part2).to eq 12
      end
    end
  end
end
