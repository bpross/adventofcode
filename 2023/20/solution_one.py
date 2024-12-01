from collections import defaultdict, deque

OFF = "OFF"
ON = "ON"
LOW = "LOW"
HIGH = "HIGH"


class FlipFlopModule:
    def __init__(self, label, destinations):
        self.label = label
        self.state = OFF
        self.destinations = destinations
        self.m_dest = defaultdict(bool)
        for dest in destinations:
            self.m_dest[dest] = True

    def receive_pulse(self, _, pulse):
        pulses = defaultdict(list)
        if pulse == HIGH:
            return pulses

        send_pulse = LOW if self.state == ON else HIGH
        self.state = ON if self.state == OFF else OFF
        for destination in self.destinations:
            pulses[send_pulse].append(destination)

        return pulses

    def has_destination(self, module):
        return self.m_dest[module]

    def __str__(self):
        return f"FlipFlopModule: {self.label} {self.state} {self.destinations} {self.m_dest}"


class ConjunctionModule:
    def __init__(self, label, destinations):
        self.label = label
        self.sources = []
        self.destinations = destinations
        self.most_recent_pulse = {}
        self.m_dest = defaultdict(bool)
        for dest in destinations:
            self.m_dest[dest] = True

    def update_sources(self, source):
        self.sources.append(source)
        self.most_recent_pulse[source.label] = LOW

    def has_destination(self, module):
        return self.m_dest[module]

    def receive_pulse(self, source, pulse):
        self.most_recent_pulse[source] = pulse
        all_high = True
        for source in self.sources:
            p = self.most_recent_pulse[source.label]
            if p == LOW:
                all_high = False

        pulses = defaultdict(list)
        pulse = HIGH
        if all_high:
            pulse = LOW

        for destination in self.destinations:
            pulses[pulse].append(destination)

        return pulses

    def __str__(self):
        return f"ConjunctionModule: {self.sources} {self.destinations}"


class Broadcaster:
    def __init__(self, destinations):
        self.label = "broadcast"
        self.destinations = destinations
        self.m_dest = defaultdict(bool)
        for dest in destinations:
            self.m_dest[dest] = True

    def receive_pulse(self, pulse):
        pulses = defaultdict(list)
        for destination in self.destinations:
            pulses[pulse].append(destination)
        return pulses

    def has_destination(self, module):
        return self.m_dest[module]

    def __str__(self):
        return f"Broadcaster: {self.destinations}"


class Untyped:
    def __init__(self, label):
        self.label = label

    def receive_pulse(self, a, b):
        return defaultdict(list)

    def has_destination(self, module):
        return False


modules = {}
flip_flops = []
conjunctions = []

with open('input.txt', 'r') as f:
    for line in f:
        module, destinations = line.strip().split(' -> ')
        destinations = destinations.split(', ')
        m = None
        label = ""
        if module.startswith('broadcaster'):
            m = Broadcaster(destinations)
            label = "broadcast"
            modules[label] = m
        elif module.startswith('%'):
            label = module[1:]
            m = FlipFlopModule(label, destinations)
            flip_flops.append(m)
            modules[label] = m
        elif module.startswith('&'):
            label = module[1:]
            m = ConjunctionModule(label, destinations)
            conjunctions.append(m)
            modules[label] = m


for c in conjunctions:
    for k, v in modules.items():
        if v.has_destination(c.label):
            c.update_sources(v)
    for n in c.destinations:
        if n not in modules:
            modules[n] = Untyped(n)

queue = deque()
button_preses = 1000
high = 0
low = 0
for i in range(button_preses):
    for d in modules['broadcast'].destinations:
        queue.append(('broadcast', d, LOW))

    low += 1 + len(modules['broadcast'].destinations)
    while len(queue) > 0:
        n = len(queue)
        for _ in range(n):
            source, module, pulse = queue.popleft()
            m = modules[module]
            pulses = m.receive_pulse(source, pulse)
            high += len(pulses[HIGH])
            low += len(pulses[LOW])
            for pulse, destinations in pulses.items():
                for d in destinations:
                    queue.append((m.label, d, pulse))


total = high * low
print(f"High: {high} Low: {low} Total: {total}")
