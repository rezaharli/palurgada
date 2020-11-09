import React from 'react';
import {View, Text} from 'react-native';
import {NavigationContainer} from '@react-navigation/native';
import {createStackNavigator} from '@react-navigation/stack';
// SCREENS
import WebViewScreen from './WebViewScreen';

const CustomScreen = () => {
  return (
    <View>
      <Text>Custom Screen</Text>
    </View>
  );
};

const Stack = createStackNavigator();

const Medium = () => {
  return (
    <NavigationContainer>
      <Stack.Navigator>
        <Stack.Screen
          name="Landing"
          component={WebViewScreen}
          initialParams={{
            url: 'https://www.google.com',
          }}
        />

        <Stack.Screen
          name="Internal"
          component={WebViewScreen}
          initialParams={{url: (props) => props.navigation.state.params.url}}
        />

        <Stack.Screen name="CustomScreen" component={CustomScreen} />
      </Stack.Navigator>
    </NavigationContainer>
  );
};

export default Medium;
