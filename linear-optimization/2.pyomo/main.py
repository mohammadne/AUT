import pyomo.environ as pyo

model = pyo.ConcreteModel()

# variables
model.x1 = pyo.Var(domain=pyo.NonNegativeReals)
model.x2 = pyo.Var(domain=pyo.NonNegativeReals)

# objective function
model.obj = pyo.Objective(expr=5*model.x1+4*model.x2, sense=pyo.maximize)

# constraints
model.const1 = pyo.Constraint(expr=model.x1+4*model.x2 <= 8)
model.const2 = pyo.Constraint(expr=-2*model.x1+model.x2 <= 4)

solvername = 'glpk'
opt = pyo.SolverFactory(solvername)
result = opt.solve(model, solvername)
model.display()

print(pyo.value(model.x1), pyo.value(model.x2))
print(pyo.value(model.obj))
