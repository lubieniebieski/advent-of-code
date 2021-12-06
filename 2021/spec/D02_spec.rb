require_relative '../solutions/D02'

describe D02 do
  context 'Real data' do
    let(:input_path) { File.join(File.dirname(__FILE__), '../inputs/D02.txt') }
    let(:input) { File.read(input_path).split("\n") }
    let(:subject) { described_class.new(input) }
    let(:part1) { subject.part1 }
    let(:part2) { subject.part2 }

    it 'Solves Part 1' do
      expect { part1 }.not_to raise_exception
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
      'forward 5
down 5
forward 8
up 3
down 8
forward 2'
    end
    let(:test_input) { test_data.split("\n") }
    let(:subject) { described_class.new(test_input) }
    context 'Part 1' do
      it 'Passes the example' do
        expect(subject.part1).to eq 150
      end
    end

    context 'Part 2' do
      it 'Passes the example' do
        expect(subject.part2).to eq 900
      end
    end
  end
end
