import math
import numpy as np
def uppgiftA(values):
    return np.mean(values)

def uppgiftB_median(values):
    median_values = []
    for val in values:
        median_values.append(abs(val - np.median(values)))
    return np.median(median_values)
def uppgiftB(values):
    mean_values = []
    for val in values:
        mean_values.append(abs(val - np.mean(values)))
    return np.mean(mean_values)
def m_variance(values):
    mean_values = []
    for val in values:
        mean_values.append((val -np.mean(values))**2)
    return np.mean(mean_values)
def m_std(values):
    m_var = m_variance(values)
    return math.sqrt(m_var)

if __name__ == "__main__":
    values = [180.0,350.0,380.0,450.0,460.0,480.0]
    answer_a = uppgiftA(values)
    answer_b = uppgiftB(values)
    answer_c_var = m_variance(values)
    answer_c_std = m_std(values)
    print str(answer_a)
    print str(answer_b)
    print str(answer_c_var),np.var(values)
    print str(answer_c_std),np.std(values)