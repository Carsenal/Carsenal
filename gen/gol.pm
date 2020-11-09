# gol.pm
# module for generating object reprsenting game state
package gol;

use strict;
use warnings;

use Game::Life;

sub print_game
{
    my $game = shift;
    print "$_\n" foreach ($game->get_text_grid());
}

sub calc_game
{
    my $args     = shift;
    my $w        = $args->{width};
    my $h        = $args->{height};
    my $disp     = $args->{disp};
    my $i_cycles = $args->{initial_cycles};
    my $s_cycles = $args->{simulate_cycles};
    my @patterns = @{ $args->{patterns} };

    my $cells = {};

    print "Simulating GOL\n";
    my $game = new Game::Life([$w, $h]);
    foreach my $pattern (@patterns) {
        $game->place_text_points($pattern->{y}, $pattern->{x}, 'O',
            @{ $pattern->{data} });
    }

    print "Simulation prepared:\n";
    print "\tDisplay: $disp\n";
    print "\twidth: $w\n";
    print "\theight: $h\n";
    print "\tinitial cycles: $i_cycles\n";
    print "\tcapture cycles: $s_cycles\n";

    if ($disp) {
        print "\n";
        print "Initial state:\n";
        print_game($game);
    }

    $game->process($i_cycles);

    print "Completed initial cell generation\n" if $disp;

    for my $i (0 .. $s_cycles) {
        $game->process;
        my @grid = $game->get_text_grid();
        $cells->{$i} = \@grid;
        if ($disp) {
            print "($i/$s_cycles):\n";
            print_game($game);
        }
    }

    print "Completed simulation\n";
    return $cells;
}

1;
