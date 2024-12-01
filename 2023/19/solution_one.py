import re


class PartsRatings:
    def __init__(self, parts_rating):
        ratings = re.search(
            r"x=(\d{1,4}),m=(\d{1,4}),a=(\d{1,4}),s=(\d{1,4})", parts_rating)
        self.x = int(ratings.group(1))
        self.m = int(ratings.group(2))
        self.a = int(ratings.group(3))
        self.s = int(ratings.group(4))

    def sum(self):
        return self.x + self.m + self.a + self.s

    def __repr__(self):
        return f"PartsRatings(x={self.x}, m={self.m}, a={self.a}, s={self.s})"


class Rule:
    def __init__(self, to_workflow, operator=None, value=None, field=None):
        self.to_workflow = to_workflow
        self.operator = operator
        self.value = value
        self.field = field

    def apply(self, parts_rating):
        if self.operator == '<':
            if getattr(parts_rating, self.field) < self.value:
                return self.to_workflow
            else:
                return None
        elif self.operator == '>':
            if getattr(parts_rating, self.field) > self.value:
                print("greater than")
                return self.to_workflow
            else:
                return None
        else:
            return self.to_workflow

    def __repr__(self):
        return f"Rule(to_workflow={self.to_workflow}, operator={self.operator}, value={self.value}, field={self.field})"


class Workflow:
    def __init__(self, workflow):
        label_and_rules = re.search(r"(\w{1,3}){(.+)}", workflow)
        self.label = label_and_rules.group(1)
        rules = label_and_rules.group(2).split(',')
        self.rules = []
        for rule in rules:
            rule_type = re.search(r"(\w{1})(<|>)(\d{1,4}):(\w{1,3})", rule)
            if rule_type:
                r = Rule(rule_type.group(4), rule_type.group(2), value=int(
                    rule_type.group(3)), field=rule_type.group(1))
                self.rules.append(r)
            else:
                r = Rule(rule)
                self.rules.append(r)

    def __repr__(self):
        return f"Workflow(label={self.label}, rules={self.rules})"


workflows = []
parts_ratings = []

workflows_parsed = False

with open('input.txt') as f:
    for line in f:
        if not workflows_parsed:
            if line == '\n':
                workflows_parsed = True
                continue
            workflows.append(line.strip())
        else:
            parts_ratings.append(line.strip())

parts_ratings = [PartsRatings(rating) for rating in parts_ratings]
workflows = [Workflow(workflow) for workflow in workflows]
w = {}
for workflow in workflows:
    w[workflow.label] = workflow

start = "in"
accepted = []
rejected = []

for rating in parts_ratings:
    print(rating)
    current = start
    not_matched = True
    while not_matched:
        print(current)
        for rule in w[current].rules:
            print("rule", rule)
            next = rule.apply(rating)
            print("next", next)
            if next == "R":
                rejected.append(rating)
                not_matched = False
                break
            elif next == "A":
                print("accepted")
                not_matched = False
                accepted.append(rating)
                break
            elif next:
                print("next", next)
                current = next
                break

total = 0
for rating in accepted:
    total += rating.sum()

print(total)
