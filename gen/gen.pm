# gen.pm
# module for generating svg for given gol
package gen;

use strict;
use warnings;

use gol;
use Template;
use Data::Dumper;

sub lpsArr
{
    my $str = shift;
    my $strlen = length $str;
    my $len = 0;
    my $i = 1;
    my @lps = (0);

    while ($i < $strlen) {
        if (substr($str, $i, 1) eq substr($str, $len, 1)) {
            $len++;
            $lps[$i] = $len;
            $i++;
        } else {
            if ($len != 0) {
                $len = $lps[$len-1];
            } else {
                $lps[$i] = 0;
                $i ++;
            }
        }
    }

    return @lps;
}

sub repeat_len
{
    my $str = shift;

    my @lps = lpsArr($str);
    my $len = $lps[(length $str) - 1];

    if ($len > ((length $str) / 2)) {
        return (length $str) - $len;
    }

    return $len;
}

sub make_opacity
{
    my $str = shift;
       $str =~ s/\./0/gm;
       $str =~ s/X/1/gm;
       $str =~ s/(.)/$1;$1;$1;/gm;
    return $str;
}

sub gen_game
{
    my ($file_args, $game_args) = @_;
    my $w = $game_args->{width};
    my $h = $game_args->{height};
    my $c = $game_args->{simulate_cycles};
    my $name = $file_args->{name};
    my $disp = $game_args->{disp};
    my $duration = $file_args->{duration};

    my $cells = gol::calc_game($game_args);
    my @res = ();

    print "Calculating per cell deltas\n";
    for my $x (0..$w) {
        for my $y (0..$h) {
            my $str = '';
            for my $i (0..$c) {
                my $row = $cells->{$i}[$y];
                next if not defined $row;
                $str .= substr($row, $x, 1);
            }
            next if (index($str, 'X') == -1);
            my $sublen = repeat_len($str);
            $sublen = length $str if $sublen == 0;
            push @res, {
                x => $x,
                y => $y,
                #period => $c,
                #status => make_opacity($str)
                period => $sublen,
                status => make_opacity(substr($str, 0, $sublen))
            };
        }
    }
    print "Initializing template\n";
    my $tt = Template->new;
    my $data = {
        pattern => $name,
        width => $w,
        height => $h,
        period => $c,
        duration => $duration,
        cells => [ @res ]
    };
    print "Printing output to $name.svg\n";
    my $output = '';
    $tt->process('gen.tt', $data, \$output);
    $output =~ s/\s+/ /gm; # Shrink all whitespace away for smaller files
    open(FH, '>', "../$name.svg") or die $!;
    print FH $output;
    close(FH);
    print "Done\n";
}

1;
