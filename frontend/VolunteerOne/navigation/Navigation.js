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
import Announcements from "../screens/Announcements";
import Explore from "../screens/Explore";
import Onboarding from "../screens/Onboarding";
import Profile from "../screens/Profile";
import Feed from "../screens/Feed";
import ViewFriends from "../screens/Profile/ViewFriends";
import ViewNotifications from "../screens/Notifications";
import Search from "../screens/Search";
import CreateAccount from "../screens/Onboarding/CreateAccount";
import Register from "../screens/Onboarding/Register";
import Login from "../screens/Onboarding/Login";
import ForgotPassword from "../screens/Onboarding/ForgotPassword";
import NewPassword from "../screens/Onboarding/NewPassword";
import Settings from "../screens/Profile/Settings";
import Pro from "../screens/Pro";
import { createBottomTabNavigator } from "@react-navigation/bottom-tabs";
import { createStackNavigator } from "@react-navigation/stack";
import ViewEvent from "../screens/Event";
import MaterialCommunityIcons from "react-native-vector-icons/MaterialCommunityIcons";
import ViewFollowing from "../screens/Profile/ViewFollowing";

/** ==================================== Routing ==================================== **/

const { width } = Dimensions.get("screen");

const Stack = createStackNavigator();
const Tab = createBottomTabNavigator();

function AnnouncementsStack(props) {
  return (
    <Stack.Navigator
      initialRouteName="Announcements"
      screenOptions={{
        gestureEnabled: false,
        mode: "card",
        headerShown: "screen",
      }}
    >
      <Stack.Screen
        name="Announcements"
        component={Announcements}
        options={{
          header: ({ navigation, scene }) => (
            <Header
              title="Announcements"
              // options
              navigation={navigation}
              scene={scene}
            />
          ),
          cardStyle: { backgroundColor: "#F8F9FE" },
        }}
      />
      <Stack.Screen
        name="Search"
        component={Search}
        options={{
          header: ({ navigation, scene }) => (
            <Header
              back
              title="Search"
              navigation={navigation}
              scene={scene}
            />
          ),
          headerTransparent: false,
          headerShown: true,
        }}
      />
      <Stack.Screen
        name="ViewNotifications"
        component={ViewNotifications}
        options={{
          header: ({ navigation, scene }) => (
            <Header
              back
              title="Notifications"
              navigation={navigation}
              scene={scene}
            />
          ),
          headerTransparent: false,
          headerShown: true,
        }}
      />
      <Stack.Screen
        name="ViewEvent"
        component={ViewEvent}
        options={{
          header: ({ navigation, scene }) => (
            <Header back title="Event" navigation={navigation} scene={scene} />
          ),
          headerTransparent: false,
          headerShown: true,
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

function FeedStack(props) {
  return (
    <Stack.Navigator
      initialRouteName="Feed"
      screenOptions={{
        gestureEnabled: false,
        mode: "card",
      }}
    >
      <Stack.Screen
        name="Feed"
        component={Feed}
        options={{
          header: ({ navigation, scene }) => (
            <Header title="Feed" navigation={navigation} scene={scene} />
          ),
          cardStyle: { backgroundColor: "#F8F9FE" },
        }}
      />
      <Stack.Screen
        name="ViewNotifications"
        component={ViewNotifications}
        options={{
          header: ({ navigation, scene }) => (
            <Header
              back
              title="Notifications"
              navigation={navigation}
              scene={scene}
            />
          ),
          headerTransparent: false,
          headerShown: true,
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
              title="Profile"
              navigation={navigation}
              scene={scene}
            />
          ),
          headerTransparent: true,
        }}
      />
      <Stack.Screen
        name="ViewFriends"
        component={ViewFriends}
        options={{
          header: ({ navigation, scene }) => (
            <Header
              back
              title="Friends"
              navigation={navigation}
              scene={scene}
            />
          ),
          headerTransparent: false,
          headerShown: true,
        }}
      />
      <Stack.Screen
        name="ViewFollowing"
        component={ViewFollowing}
        options={{
          header: ({ navigation, scene }) => (
            <Header
              back
              title="Following"
              navigation={navigation}
              scene={scene}
            />
          ),
          headerTransparent: false,
          headerShown: true,
        }}
      />
      <Stack.Screen
        name="Settings"
        component={Settings}
        options={{
          header: ({ navigation, scene }) => (
            <Header
              back
              title="Settings"
              navigation={navigation}
              scene={scene}
            />
          ),
          headerTransparent: false,
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
        name="ViewNotifications"
        component={ViewNotifications}
        options={{
          header: ({ navigation, scene }) => (
            <Header
              back
              title="Notifications"
              navigation={navigation}
              scene={scene}
            />
          ),
          headerTransparent: false,
          headerShown: true,
        }}
      />
      <Stack.Screen
        name="Explore"
        component={Explore}
        options={{
          header: ({ navigation, scene }) => (
            <Header
              // transparent
              // white
              title="Explore"
              navigation={navigation}
              scene={scene}
            />
          ),
          headerTransparent: false,
          headerShown: true,
        }}
      />
    </Stack.Navigator>
  );
}

export default function OnboardingStack(props) {
  return (
    <Stack.Navigator
      screenOptions={{
        gestureEnabled: false,
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
      initialRouteName="Announcements"
      screenOptions={{
        tabBarActiveTintColor: "#e91e63",
      }}
    >
      <Tab.Screen
        name="AnnouncementsStack"
        component={AnnouncementsStack}
        options={{
          headerShown: false,
          tabBarLabel: "Home",
          tabBarIcon: ({ color, size }) => (
            <MaterialCommunityIcons name="home" color={color} size={size} />
          ),
        }}
      />
      <Tab.Screen
        name="ExploreStack"
        component={ExploreStack}
        options={{
          headerShown: false,
          tabBarLabel: "Explore",
          tabBarIcon: ({ color, size }) => (
            <MaterialCommunityIcons name="cards" color={color} size={size} />
          ),
          // tabBarBadge: 3,
        }}
      />
      <Tab.Screen
        name="FeedStack"
        component={FeedStack}
        options={{
          headerShown: false,
          tabBarLabel: "Feed",
          tabBarIcon: ({ color, size }) => (
            <MaterialCommunityIcons
              name="account-group"
              color={color}
              size={size}
            />
          ),
          tabBarBadge: 3,
        }}
      />
      <Tab.Screen
        name="ProfileStack"
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
