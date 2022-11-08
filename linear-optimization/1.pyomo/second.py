import pyomo.environ as pyo

model = pyo.ConcreteModel()

# -------------------------------------------------------------------------------- variables

model.x111 = pyo.Var(domain=pyo.NonNegativeReals)
model.x112 = pyo.Var(domain=pyo.NonNegativeReals)
model.x121 = pyo.Var(domain=pyo.NonNegativeReals)
model.x122 = pyo.Var(domain=pyo.NonNegativeReals)
model.z = pyo.Var(domain=pyo.NonNegativeReals)
model.b1 = pyo.Var(domain=pyo.NonNegativeReals)
model.b2 = pyo.Var(domain=pyo.NonNegativeReals)
model.x211 = pyo.Var(domain=pyo.NonNegativeReals)
model.x212 = pyo.Var(domain=pyo.NonNegativeReals)
model.x221 = pyo.Var(domain=pyo.NonNegativeReals)
model.x222 = pyo.Var(domain=pyo.NonNegativeReals)

# -------------------------------------------------------------------------------- objective function

model.obj = pyo.Objective(expr=1.5*(model.x111+model.x121+model.x211+model.x221)+3*(model.x112 +
                          model.x122+model.x212+model.x222)-0.7*model.b1-1.2*model.b2-0.5*model.z, sense=pyo.maximize)

# -------------------------------------------------------------------------------- constraints

# initial values for first month
model.const1 = pyo.Constraint(expr=model.x111+model.x112 == 2000)
model.const2 = pyo.Constraint(expr=model.x121+model.x122 == 2000)

# first month sales
model.const3 = pyo.Constraint(expr=model.x111+model.x121 <= 2000)
model.const4 = pyo.Constraint(expr=model.x112+model.x122 <= 1000)

# transfer
model.const5 = pyo.Constraint(expr=model.z <= 2000 - model.x121 - model.x122)

# initial values for second month
model.const6 = pyo.Constraint(expr=model.x211+model.x212 == model.b1+model.z)
model.const7 = pyo.Constraint(expr=model.x221+model.x222 == model.b2)

# second month sales
model.const8 = pyo.Constraint(expr=model.x211+model.x221 <= 2000)
model.const9 = pyo.Constraint(expr=model.x212+model.x222 <= 1000)

# ratios
model.const10 = pyo.Constraint(expr=4*model.x121 >= model.x111)
model.const11 = pyo.Constraint(expr=2*model.x122 >= 3*model.x112)
model.const12 = pyo.Constraint(expr=4*model.x221 >= model.x211)
model.const13 = pyo.Constraint(expr=2*model.x222 >= 3*model.x212)

solvername = 'glpk'
opt = pyo.SolverFactory(solvername)
result = opt.solve(model, solvername)

model.display()
