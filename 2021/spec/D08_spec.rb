require_relative '../solutions/D08'
describe D08::Entry do
  let(:entry) do
    described_class.new('acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf')
  end
  it { expect(entry.decode_digit('dab')).to eq 7 }
  it { expect(entry.decode_digit('ab')).to eq 1 }
  it { expect(entry.decode_digit('acedgfb')).to eq 8 }

  it { expect(entry.decode_digit('cdfbe')).to eq 5 }
  it { expect(entry.decode_digit('gcdfa')).to eq 2 }
  it { expect(entry.decode_digit('fbcad')).to eq 3 }
  it { expect(entry.decode_digit('cefabd')).to eq 9 }
  it { expect(entry.decode_digit('cdfgeb')).to eq 6 }
  it 'returns the right output' do
    expect(entry.output).to eq 5353
  end

  let(:another_entry) do
    described_class.new('edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec |   fcgedb cgb dgebacf gc')
  end
  it { expect(another_entry.decode_digit('fcgedb')).to eq 9 }
  it { expect(another_entry.decode_digit('cgb')).to eq 7 }
end
describe D08 do
  context 'Real data' do
    let(:input_path) { File.join(File.dirname(__FILE__), '../inputs/D08.txt') }
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
      'be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
      edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
      fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
      fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
      aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
      fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
      dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
      bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
      egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
      gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce'
    end
    let(:test_input) { test_data.split("\n") }
    let(:subject) { described_class.new(test_input) }
    context 'Part 1' do
      it 'Passes the example' do
        expect(subject.part1).to eq 26
      end
    end

    context 'Part 2' do
      it 'Passes the example' do
        expect(subject.part2).to eq 61_229
      end
    end
  end
end
