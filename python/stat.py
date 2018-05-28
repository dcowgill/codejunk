import numpy

def weighttest(x, w):
    if w is None:
        try:
            from ErrorVal import NegErrs, PosErrs, PrimeVals
            w = 1/(NegErrs(x)+PosErrs(x))
            y = PrimeVals(x)
        except (ImportError, AttributeError):
            w = numpy.ones_like(x)
            y = x
    else:
        try:
            from ErrorVal import PrimeVals
            y = PrimeVals(x)
        except (ImportError, AttributeError):
            y = x
    w = w*numpy.abs(numpy.isnan(y)-1) #makes sure that nan data points don't have a weight
    return numpy.array(y), w

def mean(x, w=None, axis=None, NN=True):
    x, w = weighttest(x, w)
    if NN:
        result = 1.*numpy.sum(x*w, axis=axis)/numpy.sum(w, axis=axis)
    else:
        result = 1.*numpy.nansum(x*w, axis=axis)/numpy.nansum(w, axis=axis)
    return result

def variance(x, w=None, axis=None, NN=True):
    x, w = weighttest(x, w)
    if NN:
        result = sumsqrdev(x, w, axis, NN)/(numpy.sum(w)-1)
    else:
        result = sumsqrdev(x, w, axis, NN)/(numpy.nansum(w)-1)
    return result

def pvariance(x, w=None, axis=None, NN=True):
    x, w = weighttest(x, w)
    if NN:
        result = sumsqrdev(x, w, axis, NN)/numpy.sum(w, axis=axis)
    else:
        result = sumsqrdev(x, w, axis, NN)/numpy.nansum(w, axis=axis)
    return result

def sumsqrdev(x, w=None, axis=None, NN=True):
    result = moment(x, 2, w, axis, NN)
    return result

def moment(x, n, w=None, axis=None, NN=True):
    x, w = weighttest(x, w)
    if NN:
        result = numpy.sum(w*(x - mean(x, w, axis, NN))**n, axis=axis)
    else:
        result = numpy.nansum(w*(x - mean(x, w, axis, NN))**n, axis=axis)
    return result
