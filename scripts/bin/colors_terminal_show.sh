perl -CADS -lwe '
my $block = shift || (chr(0x2588) x 3);
for (["", 0], ["1;", 0], ["", 8], ["1;", 8]) {
  my ($bold, $offset) = @$_;
  my @range = map $offset + $_, 0..7;
  printf "%s %-6s ", $bold ? "bold" : "norm", "$range[0]-$range[-1]";
  print map("\e[${bold}38;5;${_}m$block", @range), "\e[0m"
}
' "$@"
