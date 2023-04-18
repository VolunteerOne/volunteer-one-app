import React from "react";
import {
  StyleSheet,
  Text,
  View,
  Dimensions,
  ImageBackground,
  Animated,
  PanResponder,
  TouchableOpacity
} from "react-native";

const ExploreImage = ({ item }) => {
    return (
    <ImageBackground
        style={{
        flex: 1,
        height: null,
        width: null,
        resizeMode: "cover",
        borderRadius: 20,
        }}
        source={item.uri} >
        <View style = {styles.overlay}>
        <Text style = {styles.titleText}>{item.title}</Text>
        <Text style = {styles.eventDetailsText}>{item.author} </Text>
        <Text style = {styles.postedText}>posted {item.date}</Text>
        <Text style = {styles.causeAreas}> Cause Areas: </Text>
        <Text style = {styles.causeAreas2}>{item.causeAreas}</Text>
        <TouchableOpacity><Text style = {styles.button}>More details</Text></TouchableOpacity>
        </View>
    </ImageBackground>
    );
  };
  
  const styles = StyleSheet.create({
    overlay: {
        marginTop: 475,
        backgroundColor: "#FFFFFF",
        height: 175,
        opacity: 0.8,
        borderRadius: 5
    },
    titleText: {
        fontSize: 24,
        fontWeight: "bold",
        color: "#32325D",
        marginTop: 10,
        marginLeft: 15
      },
      eventDetailsText: {
        fontSize: 16,
        color: "#32325D",
        marginTop: 10,
        marginLeft: 15
      },
      postedText: {
        fontSize: 16,
        color: "#32325D",
        marginTop: 2.5,
        marginLeft: 15
      },
    shadowProp: {
      shadowColor: "#171717",
      shadowOffset: { width: -2, height: 4 },
      shadowOpacity: 0.2,
      shadowRadius: 3,
    },
    causeAreas: {
        fontSize: 14,
        color: "#000000",
        marginTop: 10,
        marginLeft: 10,
        fontWeight: "bold",

    },
    causeAreas2: {
        fontSize: 14,
        color: "#32325D",
        marginLeft: 15,

    },
    button: {
        fontWeight: "bold",
        fontSize: 13,
        color: "#5E72E4",
        marginLeft: 15,
        marginTop: 8,
        textDecorationLine: 'underline'
    }
  });
  
  export default ExploreImage;
  


