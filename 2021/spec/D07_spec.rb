require_relative '../solutions/D07'

describe D07 do
  context 'Real data' do
    let(:input_path) { File.join(File.dirname(__FILE__), '../inputs/D07.txt') }
    let(:input) { File.read(input_path) }
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
      '16,1,2,0,4,2,7,1,2,14'
    end
    let(:test_input) { test_data }
    let(:subject) { described_class.new(test_input) }
    context 'Part 1' do
      it 'Passes the example' do
        expect(subject.part1).to eq 37
      end
    end

    context 'Part 2' do
      it 'Passes the example' do
        expect(subject.part2).to eq 168
      end
    end
  end
end
