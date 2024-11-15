<?php
include('vendor/autoload.php');

/**
 * This is project's console commands configuration for Robo task runner.
 *
 * @see https://robo.li/
 */
class RoboFile extends \Robo\Tasks
{
    public function lang()
    {
        $domain = 'admin.pocketstore.io';
        $adminUrl = 'https://' . $domain;

        exec('mkdir -p i18n/locales');

        $pb = new \Pb\Client($adminUrl);
        $languages = ['de', 'en'];

        foreach ($languages as $language) {
            $file[$language] = [];
            $response = $pb->collection('translations')->getFullList(['filter' => 'lang="' . $language . '"']);

            if (!empty($response['items']) && count($response['items']) > 0) {
                foreach ($response['items'] as $item) {
                    $explode = explode('.', $item['key']);

                    if (count($explode) === 2) {                       
                        if(!empty($file[$language][$explode[0]][$explode[1]])){
                            $file[$language][$explode[0]][$explode[1]] =  [];
                        }
                        else {
                            $file[$language][$explode[0]][$explode[1]] = [$explode[1] => $item['translated']];
                        }
                    } else if (count($explode) === 3) {                        
                        if(empty($file[$language][$explode[0]][$explode[1]][$explode[2]])){
                            $file[$language][$explode[0]][$explode[1]][$explode[2]] = [$explode[2] => $item['translated']];
                        }
                        else {
                            $file[$language][$explode[0]][$explode[1]][$explode[2]] = array_merge(
                                [$explode[2] => $item['translated']], 
                                $file[$language][$explode[0]][$explode[1]][$explode[2]]
                            );
                        }
                    }
                }
            }

        }


        foreach ($languages as $language) {
            $file[$language] = $this->recursiveTransform($file[$language]);
            file_put_contents('i18n/locales/' . $language . '.json', json_encode($file[$language], JSON_FORCE_OBJECT | JSON_UNESCAPED_UNICODE));
        }
    }

    public function recursiveTransform(array $file)
    {
        foreach ($file as $key => $item) {
            // If the item is an array with exactly one element, replace it with the element itself
            if (is_array($file[$key]) && count($file[$key]) === 1) {
                $file[$key] = $file[$key][array_key_first($file[$key])];
            } // If it's an array with multiple elements, recursively process it
            else if (is_array($file[$key])) {
                $file[$key] = $this->recursiveTransform($file[$key]);
            }
        }

        return $file;
    }

}