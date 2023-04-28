import React from 'react';
import { withNavigation } from '@react-navigation/compat';
import PropTypes from 'prop-types';
import { StyleSheet, Dimensions, Image, TouchableWithoutFeedback } from 'react-native';
import { Block, Text, theme } from 'galio-framework';
import { argonTheme } from '../constants';

class UserAndOrgItem extends React.Component {
  render() {
    const { navigation, item, style } = this.props;

    return (
      <TouchableWithoutFeedback onPress={() => navigation.navigate('Profile', {theUser: item.name,})}>
        <Block row card style={[styles.card, styles.shadow, style]}>
          <Block flex={1} style={styles.imageContainer}>
            <Image source={{uri: item.image}} style={styles.image} />
          </Block>
          <Block flex={3} style={styles.contentContainer}>
            <Text size={16} style={styles.cardTitle}>{item.name}</Text>
          </Block>
        </Block>
      </TouchableWithoutFeedback>
    );
  }
}

UserAndOrgItem.propTypes = {
  item: PropTypes.object,
  style: PropTypes.any,
}

const styles = StyleSheet.create({
  card: {
    backgroundColor: theme.COLORS.WHITE,
    marginVertical: theme.SIZES.BASE / 4,
    minHeight: 45,
    alignItems: 'center',
    justifyContent: 'center',
  },
  cardTitle: {
    flexWrap: 'wrap',
    paddingLeft: theme.SIZES.BASE / 2,
    textAlign: 'left',
  },
  imageContainer: {
    justifyContent: 'center',
    alignItems: 'flex-end',
    paddingRight: theme.SIZES.BASE / 2,
  },
  contentContainer: {
    justifyContent: 'center',
  },
  image: {
    height: 50,
    width: 50,
    borderRadius: 30,
    marginTop: 3,
    marginBottom: 3,
  },
  shadow: {
    shadowColor: theme.COLORS.BLACK,
    shadowOffset: { width: 0, height: 2 },
    shadowRadius: 4,
    shadowOpacity: 0.1,
    elevation: 2,
  },
});

export default withNavigation(UserAndOrgItem);
