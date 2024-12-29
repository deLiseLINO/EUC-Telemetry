CREATE TABLE default.metrics
(
    -- Date and Time
    date Date,
    time DateTime,

    -- GPS Data
    latitude Float64,
    longitude Float64,
    gps_speed Float32,
    gps_alt Float32,
    gps_heading Float32,
    gps_distance Int32,

    -- Speed and Power
    speed Float32,
    voltage Float32,
    phase_current Float32,
    current Float32,
    power Float32,
    torque Float32,
    pwm Float32,

    -- Battery and Distance
    battery_level Int32,
    distance Int32,
    totaldistance Int32,

    -- Temperature
    system_temp Int32,
    temp2 Int32,

    -- Orientation
    tilt Float32,
    roll Float32,

    -- Mode and Alerts
    mode Int32,
    alert Int32
)
ENGINE = MergeTree
ORDER BY (date, time)
