import React from "react";
import { Animated, Dimensions, Easing, TouchableOpacity } from "react-native";
// header for screens
import { Header, Icon } from "../components";
import { argonTheme, tabs } from "../constants";

import Articles from "../screens/Articles";
import { Block, theme } from "galio-framework";
// drawer
// import CustomDrawerContent from "./Menu";
import Elements from "../screens/Elements";
// screens
import Home from "../screens/Home";
import Onboarding from "../screens/Onboarding";
import Profile from "../screens/Profile";
import Friends from "../screens/Friends";
import CreateAccount from "../screens/Onboarding/CreateAccount";
import Register from "../screens/Onboarding/Register";
import Login from "../screens/Onboarding/Login";
import ForgotPassword from "../screens/Onboarding/ForgotPassword";
import NewPassword from "../screens/Onboarding/NewPassword";
import Pro from "../screens/Pro";
import { createBottomTabNavigator } from "@react-navigation/bottom-tabs";
import { createStackNavigator } from "@react-navigation/stack";

import MaterialCommunityIcons from "react-native-vector-icons/MaterialCommunityIcons";

/** ==================================== Routing ==================================== **/

const { width } = Dimensions.get("screen");

const Stack = createStackNavigator();
const Tab = createBottomTabNavigator();

function HomeStack(props) {
  return (
    <Stack.Navigator
      screenOptions={{
        mode: "card",
        headerShown: "screen",
        gestureEnabled: false,
      }}
    >
      <Stack.Screen
        name="Home"
        component={Home}
        options={{
          header: ({ navigation, scene }) => (
            <Header
              title="Home"
              options
              navigation={navigation}
              scene={scene}
            />
          ),
          cardStyle: { backgroundColor: "#F8F9FE" },
        }}
      />
      <Stack.Screen
        name="Pro"
        component={Pro}
        options={{
          header: ({ navigation, scene }) => (
            <Header
              back
              transparent
              white
              title=""
              navigation={navigation}
              scene={scene}
            />
          ),
          headerTransparent: true,
          headerShown: true,
        }}
      />
    </Stack.Navigator>
  );
}

function FriendsStack(props) {
  return (
    <Stack.Navigator
      screenOptions={{
        mode: "card",
      }}
    >
      <Stack.Screen
        name="Friends"
        component={Friends}
        options={{
          header: ({ navigation, scene }) => (
            <Header title="Friends" navigation={navigation} scene={scene} />
          ),
          cardStyle: { backgroundColor: "#F8F9FE" },
        }}
      />
      <Stack.Screen
        name="Pro"
        component={Pro}
        options={{
          header: ({ navigation, scene }) => (
            <Header
              back
              transparent
              white
              title=""
              navigation={navigation}
              scene={scene}
            />
          ),
          headerTransparent: true,
          headerShown: true,
        }}
      />
    </Stack.Navigator>
  );
}

function ProfileStack(props) {
  return (
    <Stack.Navigator
      initialRouteName="Profile"
      screenOptions={{
        mode: "card",
        headerShown: "screen",
      }}
    >
      <Stack.Screen
        name="Profile"
        component={Profile}
        options={{
          header: ({ navigation, scene }) => (
            <Header
              transparent
              white
              title=""
              navigation={navigation}
              scene={scene}
            />
          ),
          headerTransparent: true,
        }}
      />
      <Stack.Screen
        name="Pro"
        component={Pro}
        options={{
          header: ({ navigation, scene }) => (
            <Header
              back
              transparent
              white
              title=""
              navigation={navigation}
              scene={scene}
            />
          ),
          headerTransparent: true,
          headerShown: true,
        }}
      />
    </Stack.Navigator>
  );
}

function ExploreStack(props) {
  return (
    <Stack.Navigator
      initialRouteName="Explore"
      screenOptions={{
        mode: "card",
        headerShown: "screen",
      }}
    >
      <Stack.Screen
        name="Explore"
        component={Pro}
        options={{
          header: ({ navigation, scene }) => (
            <Header
              back
              transparent
              white
              title="Explore"
              navigation={navigation}
              scene={scene}
            />
          ),
          headerTransparent: true,
        }}
      />
    </Stack.Navigator>
  );
}

export default function OnboardingStack(props) {
  return (
    <Stack.Navigator
      screenOptions={{
        mode: "card",
        headerShown: false,
        gestureEnabled: false,
      }}
    >
      <Stack.Screen
        name="Onboarding"
        component={Onboarding}
        option={{
          headerTransparent: true,
        }}
      />
      <Stack.Screen
        name="Login"
        component={Login}
        options={{
          header: ({ navigation, scene }) => (
            <Header
              transparent
              white
              title="Login"
              navigation={navigation}
              scene={scene}
            />
          ),
          headerTransparent: true,
          headerShown: true,
        }}
      />
      <Stack.Screen
        name="CreateAccount"
        component={CreateAccount}
        options={{
          header: ({ navigation, scene }) => (
            <Header
              back
              transparent
              white
              title="Create new account"
              navigation={navigation}
              scene={scene}
            />
          ),
          headerTransparent: true,
          headerShown: true,
        }}
      />
      <Stack.Screen
        name="Register"
        component={Register}
        options={{
          header: ({ navigation, scene }) => (
            <Header
              back
              transparent
              white
              title="Register"
              navigation={navigation}
              scene={scene}
            />
          ),
          headerTransparent: true,
          headerShown: true,
        }}
      />
      <Stack.Screen
        name="ForgotPassword"
        component={ForgotPassword}
        options={{
          header: ({ navigation, scene }) => (
            <Header
              back
              transparent
              white
              title="Forgot Password"
              navigation={navigation}
              scene={scene}
            />
          ),
          headerTransparent: true,
          headerShown: true,
        }}
      />
      <Stack.Screen
        name="NewPassword"
        component={NewPassword}
        options={{
          header: ({ navigation, scene }) => (
            <Header
              transparent
              white
              title="Choose new password"
              navigation={navigation}
              scene={scene}
            />
          ),
          headerTransparent: true,
          headerShown: true,
        }}
      />
      <Stack.Screen name="App" component={BottomNav} />
    </Stack.Navigator>
  );
}

function BottomNav() {
  return (
    <Tab.Navigator
      initialRouteName="Home"
      screenOptions={{
        tabBarActiveTintColor: "#e91e63",
      }}
    >
      <Tab.Screen
        name="Home"
        component={HomeStack}
        options={{
          headerShown: false,
          tabBarLabel: "Home",
          tabBarIcon: ({ color, size }) => (
            <MaterialCommunityIcons name="home" color={color} size={size} />
          ),
        }}
      />
      <Tab.Screen
        name="Explore"
        component={ExploreStack}
        options={{
          headerShown: false,
          tabBarLabel: "Explore",
          tabBarIcon: ({ color, size }) => (
            <MaterialCommunityIcons name="bell" color={color} size={size} />
          ),
          tabBarBadge: 3,
        }}
      />
      <Tab.Screen
        name="Friends"
        component={FriendsStack}
        options={{
          headerShown: false,
          tabBarLabel: "Friends",
          tabBarIcon: ({ color, size }) => (
            <MaterialCommunityIcons name="account-group" color={color} size={size} />
          ),
          tabBarBadge: 3,
        }}
      />
      <Tab.Screen
        name="Profile"
        component={ProfileStack}
        options={{
          headerShown: false,
          tabBarLabel: "Profile",
          tabBarIcon: ({ color, size }) => (
            <MaterialCommunityIcons name="account" color={color} size={size} />
          ),
        }}
      />
    </Tab.Navigator>
  );
}
