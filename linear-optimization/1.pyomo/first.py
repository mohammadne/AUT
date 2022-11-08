import pyomo.environ as pyo

model = pyo.ConcreteModel()

# variables
model.x1 = pyo.Var(domain=pyo.NonNegativeIntegers)
model.x2 = pyo.Var(domain=pyo.NonNegativeIntegers)
model.y1 = pyo.Var(domain=pyo.NonNegativeIntegers)
model.y2 = pyo.Var(domain=pyo.NonNegativeIntegers)

# objective function
model.obj = pyo.Objective(expr=100*model.x1+30 *
                          model.x2+80*model.y1+30*model.y2, sense=pyo.maximize)

# constraints
model.const1 = pyo.Constraint(
    expr=40*(model.x1+model.x2)+30*(model.y1+model.y2) <= 40000)
model.const2 = pyo.Constraint(
    expr=5*model.x1+2*model.x2+4.5*model.y1+2.5*model.y2 <= 6000)

solvername = 'glpk'
opt = pyo.SolverFactory(solvername)
result = opt.solve(model, solvername)
model.display()

print(pyo.value(model.x1), pyo.value(model.x2),
      pyo.value(model.y1), pyo.value(model.y2))
print(pyo.value(model.obj))
