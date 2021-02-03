<?php
// source code: http://www.php-benchmark-script.com
// Modified to PHP 8.0 and add Fibonacci test by Dinda <dindasigma@gmail.com>


set_time_limit(60);

function test_Math($count = 1_400_000) {
    $time_start = microtime(true);
    $mathFunctions = ["abs", "sqrt", "acos", "asin", "atan", "floor", "exp", "sin", "tan", "is_finite", "is_nan"];
    for ($i=0; $i < $count; $i++) {
        foreach ($mathFunctions as $function) {
            call_user_func_array($function, [$i]);
        }
    }
    return number_format(microtime(true) - $time_start, 3);
}

function test_Loops($count = 190_000_000) {
    $time_start = microtime(true);
    for($i = 0; $i < $count; ++$i);
    $i = 0; while($i < $count) ++$i;
    return number_format(microtime(true) - $time_start, 3);
}

function test_Fibonacci($count = 35) {
    $time_start = microtime(true);
    $c = $count < 2 ? 1 : test_fibonacci(count: $count -2) + test_fibonacci(count: $count - 1);
    return number_format(microtime(true) - $time_start, 3);
}

function test_IfElse($count = 90_000_000) {
    $time_start = microtime(true);
    for ($i=0; $i < $count; $i++) {
        if ($i == -1) {
        } elseif ($i == -2) {
        } else if ($i == -3) {
        }
    }
    return number_format(microtime(true) - $time_start, 3);
}

function test_StringManipulation($count = 1_300_000) {
    $time_start = microtime(true);
    $stringFunctions = ["addslashes", "chunk_split", "metaphone", "strip_tags", "md5", "sha1", "strtoupper", "strtolower", "strrev", "strlen", "soundex", "ord"];
    foreach ($stringFunctions as $key => $function) {
        if (!function_exists($function)) unset($stringFunctions[$key]);
    }
    $string = "the quick brown fox jumps over the lazy dog";
    for ($i=0; $i < $count; $i++) {
        foreach ($stringFunctions as $function) {
            $r = call_user_func_array($function, [$string]);
        }
    }
    return number_format(microtime(true) - $time_start, 3);
}


$total = 0;
$functions = get_defined_functions();

$line = str_pad("-",38,"-");
$title = str_pad("JIT PHP 8.0 BENCHMARK",36," ",STR_PAD_BOTH);
echo "<pre>";
echo sprintf("%s\n|%s|\n%s\n", $line, $title, $line);
echo sprintf("Start: %s\nServer: %s\nPHP version : %s\nPlatform: %s\n", date("Y-m-d"), $_SERVER['SERVER_ADDR'], PHP_VERSION, PHP_OS);

$jit = function_exists('opcache_get_status') ? opcache_get_status()['jit'] : false;


$opt_level = [
    4 => 'Tracing',
    5 => 'Function',
];
if ($jit) {
    $jit_buffer_size = sprintf("%d MB", number_format($jit['buffer_size'] / 1_048_576, 0));
    echo sprintf("JIT Enable: %s\nBuffer Size: %s\nFlags: %s\n", $jit['enabled'], $jit_buffer_size, $opt_level[$jit['opt_level']]);
}
else {
    echo sprintf("JIT Enable: %d\n", 0);
}

echo $line."\n";

foreach ($functions['user'] as $user) {
    if (preg_match('/^test_/', $user)) {
        $total += $result = $user();
        echo sprintf("%s : %f sec\n", str_pad($user, 23, " "), $result);
    }
}

echo sprintf("%s\n%s : %f sec\n", $line, str_pad("Total time", 23, " "), $total);
echo "</pre>";