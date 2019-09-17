#!/usr/bin/python

import itertools
import math
import trueskill

#  You can think of beta as the number of skill points to guarantee about an 80%
#  chance of winning.
BETA = 60

def win_probability(team1, team2):
  delta_mu = sum(r.mu for r in team1) - sum(r.mu for r in team2)
  sum_sigma = sum(r.sigma ** 2 for r in itertools.chain(team1, team2))
  size = len(team1) + len(team2)
  denom = math.sqrt(size * (BETA * BETA) + sum_sigma)
  ts = trueskill.global_env()
  return ts.cdf(delta_mu / denom)

if __name__ == "__main__":
  r1 = trueskill.Rating(mu=80)
  r2 = trueskill.Rating(mu=80)

  for x in xrange(20):
    r1, r2 = trueskill.rate_1vs1(r1, r2)
    print x+1, r1, r2
    print win_probability([r1], [r2])
