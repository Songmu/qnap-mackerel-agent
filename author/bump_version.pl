#!/usr/bin/env perl
use 5.014;
use warnings;
use utf8;
use autodie;

my ($changelog, $qpkg_conf) = @ARGV;

my $version_reg = qr/[0-9]+(?:\.[0-9]+){2}/;
my $date_reg = qr/20[0-9]{2}-[0-9]{2}-[0-9]{2}/;

my ($version, $date);
open my $fh, '<', $changelog;
while (my $l = <$fh>) {
    if ($l =~ /^## ($version_reg) \(($date_reg)\)/) {
        ($version, $date) = ($1, $2);
        last;
    }
}

my $content = do {
    local $/;
    open my $fh, '<', $qpkg_conf;
    <$fh>;
};

$content =~ s/$version_reg/$version/;
$content =~ s/$date_reg/$date/;

open my $fh2, '>', $qpkg_conf;
print $fh2 $content;
