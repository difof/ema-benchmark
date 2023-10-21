
export function ema(priceAddr: i32, len: i32, period: i32): i32 {
    let resultAddr: i32 = priceAddr + len;
    let numPrices: i32 = len / 8;
    let k: f64 = 2.0 / (period + 1.0);
    writeFloat(resultAddr, readFloat(priceAddr));

    for (let i: i32 = 1; i < numPrices; i++) {
        // ema[i] = (data[i] * k) + (ema[i-1] * (1 - k))
        writeFloat(resultAddr + i * 8,
            (readFloat(priceAddr + i * 8) * k) +
            (readFloat(resultAddr + (i - 1) * 8) * (1 - k))
        );
    }

    return resultAddr;
}

function readFloat(addr: i32): f64 {
    return load<f64>(addr);
}

function writeFloat(addr: i32, value: f64): void {
    store<f64>(addr, value);
}
