require "option_parser"

distributions = [] of String

parser = OptionParser.parse do |parser|
  parser.banner = "Usage: distrohop [arguments]"
  parser.on("-v", "--version", "Show version") {puts "version 1.0"}
  parser.on("-l LIST", "--list=LIST", "Example: $ distrohopper -l \"Ubuntu Debian Mint\"") { |list| distributions = list.split }
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

def calc_winner(distros)
  rng = Random.new
  winner = distros[rng.rand(distros.size)]
  puts "You should install #{winner}"
end

if distributions.size > 0
  calc_winner(distributions)
else
  puts parser
  exit(1)
end