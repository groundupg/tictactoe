# 27.10.2025

The primary utility of a computer is it's capacity for mass computation. It is this capacity for
large computation which saw the computer replacing humans as the defacto means of computation.
This utility remains the primary use of the computer today.

What does not require large computation may not require a computer; it may be argued that the task of defining a program which must be *computable* on a machine is a layer which, if possible, one
should avoid.

However, we should not disregard the computer; the utility of large computation is in fact one very
useful -- one which may enable the human to reach heights never before achievable.

What I propose is the utilisation of the computer's ability to perform large computations.

## On Decision Making

*A decision can be weighted by the expected value of making the decision.*

*The expected value of a decision is the weighted sum of the multiplication of the value and
probability of all possible outcomes**

```
  type Outcome:
      probability: float
      value: float

  fn o_ev (p: float v: float) -> float:
      return p * v

  fn EV (no: list[Outcome]) -> float:
      current: float
      for o in no:
        current += o_ev(o.probability, o.value)
      return current 
```

## Simulation

The above are very simple computations to calculate the EV of a given decision. However, the computation
has a requirement -- the list of possible outcomes of making a decision, with their probability
& value.

It is precisely these areas where humans struggle, in establishing the possible outcomes of taking
a decision. Very simple games can be handled, i.e. checkers, but when more factors are introduced
humans crumble.

I propose that the most valuable way we, as humans, can utilise the computer, is through simulation.
Specifically, simulation of possible outcomes of a decision which enable a human to make expected
value decisions.

This has proven to be extremely beneficial in many cases: chess, poker, go. The afore mentioned
are games with clear constraints that can be defined, enabling the mass simulation of outcomes
which facilitate expected value decisions for any single decision.

I, however, think that we can do more.
I think that we can use the simulation capacity of computers to apply these methods to business &
science.

### A Simulation Function

A *Simulator* works through the application of a simulate function applied to a decision n times.
Therefore, a *simulate* function must be defined.

Take a business MS.
MS has a baseline constraint of cash flow, £8500.
The value metric of MS is the total amount of capital over 30 years.

MS wants a business strategy which maximises their value metric while never dropping below their
baseline.

```
INVESTMENTS: float
PROFIT: float
EXPENSES: float
BASELINE: float = 8500

fn capital (i, p):
    return i + p


fn simulate (d: Decision, i: float, p: float) -> Outcome:
    c := capital(i, p)
    for i:= 0; i < 30; i++:
        
```
