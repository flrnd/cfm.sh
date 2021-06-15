require "option_parser"

options_list = [] of String

parser = OptionParser.parse do |parser|
  parser.banner = "Usage: cfm.sh [arguments]"
  parser.on("-v", "--version", "Show version") {puts "version 1.0"}
  parser.on("-l LIST", "--list=LIST", "Example: $ cfm.sh -l \"Ubuntu Debian Mint\"") { |list| options_list = list.split }
  parser.on("-h", "--help", "Show help") do
    puts parser
    exit
  end
  parser.invalid_option do |flag|
    STDERR.puts "ERROR: #{flag} is not a valid option"
    STDERR.puts parser
    exit(1)
  end
end

def calc_winner(from_list)
  rng = Random.new
  winner = from_list[rng.rand(from_list.size)]
  puts " #{winner}"
end

if options_list.size > 0
  calc_winner(options_list)
else
  puts parser
  exit(1)
end
